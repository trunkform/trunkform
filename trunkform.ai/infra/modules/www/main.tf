variable "domain_name" {
  type = string
}

variable "environment" {
  type = string
}

variable "account" {
  type = string
}

variable "region" {
  type = string
}

variable "zone_id" {
  type = string
}

locals {
  bucket  = "${replace(var.domain_name, ".", "-")}-${var.account}"
  region  = var.region
  zone_id = var.zone_id

}

resource "aws_s3_bucket" "www" {
  bucket              = local.bucket
  force_destroy       = null
  object_lock_enabled = false
  tags                = {}
}

resource "aws_s3_bucket_public_access_block" "www" {
  bucket                  = local.bucket
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true

  depends_on = [aws_s3_bucket.www]
}

resource "aws_s3_bucket_server_side_encryption_configuration" "www" {
  bucket = local.bucket
  rule {
    bucket_key_enabled = true
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }

  depends_on = [aws_s3_bucket.www]
}

resource "aws_s3_bucket_versioning" "www" {
  bucket = local.bucket
  versioning_configuration {
    status = "Disabled"
  }

  depends_on = [aws_s3_bucket.www]
}

resource "aws_cloudfront_origin_access_control" "www" {
  name                              = "${local.bucket}.s3.${local.region}.amazonaws.com"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

resource "aws_acm_certificate" "www" {
  provider = aws.global

  domain_name               = var.domain_name
  key_algorithm             = "RSA_2048"
  subject_alternative_names = [var.domain_name]
  validation_method         = "DNS"
  tags                      = {}

  options {
    certificate_transparency_logging_preference = "ENABLED"
  }
}

locals {
  acm_validation = [for d in aws_acm_certificate.www.domain_validation_options : d if d.domain_name == var.domain_name][0]
}

resource "aws_route53_record" "cname" {
  provider = aws.global

  name    = local.acm_validation.resource_record_name
  records = [local.acm_validation.resource_record_value]
  ttl     = 3600
  type    = "CNAME"
  zone_id = local.zone_id
}

resource "aws_acm_certificate_validation" "www" {
  provider = aws.global

  certificate_arn         = aws_acm_certificate.www.arn
  validation_record_fqdns = [aws_route53_record.cname.fqdn]
}

resource "aws_cloudfront_distribution" "www" {
  aliases             = [var.domain_name]
  default_root_object = "index.html"
  enabled             = true
  http_version        = "http2"
  is_ipv6_enabled     = false
  price_class         = "PriceClass_100"
  tags                = {}

  default_cache_behavior {
    allowed_methods            = ["GET", "HEAD"]
    cache_policy_id            = "658327ea-f89d-4fab-a63d-7e88639e58f6" # CachingOptimized
    cached_methods             = ["GET", "HEAD"]
    compress                   = true
    default_ttl                = 0
    max_ttl                    = 0
    min_ttl                    = 0
    origin_request_policy_id   = "88a5eaf4-2fd4-4709-b370-b4c650ea3fcf" # CORS-S3Origin
    response_headers_policy_id = "60669652-455b-4ae9-85a4-c4c02393f86c" # SimpleCORS
    smooth_streaming           = false
    target_origin_id           = "${local.bucket}.s3.${local.region}.amazonaws.com"
    viewer_protocol_policy     = "redirect-to-https"

    grpc_config {
      enabled = false
    }
  }

  origin {
    connection_attempts      = 3
    connection_timeout       = 10
    domain_name              = "${local.bucket}.s3.${local.region}.amazonaws.com"
    origin_access_control_id = aws_cloudfront_origin_access_control.www.id
    origin_id                = "${local.bucket}.s3.${local.region}.amazonaws.com"
  }

  custom_error_response {
    error_code         = 403
    response_code      = 200
    response_page_path = "/index.html"
  }

  custom_error_response {
    error_code         = 404
    response_code      = 200
    response_page_path = "/index.html"
  }

  restrictions {
    geo_restriction {
      locations        = []
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn      = aws_acm_certificate.www.arn
    minimum_protocol_version = "TLSv1.2_2021"
    ssl_support_method       = "sni-only"
  }

  depends_on = [aws_acm_certificate_validation.www]
}

resource "aws_route53_record" "a" {
  provider = aws.global

  name    = var.domain_name
  type    = "A"
  zone_id = local.zone_id

  alias {
    name                   = aws_cloudfront_distribution.www.domain_name
    zone_id                = aws_cloudfront_distribution.www.hosted_zone_id
    evaluate_target_health = false
  }
}

output "distribution_id" {
  value = aws_cloudfront_distribution.www.id
}

output "distribution_domain_name" {
  value = aws_cloudfront_distribution.www.domain_name
}

output "distribution_zone_id" {
  value = aws_cloudfront_distribution.www.hosted_zone_id
}

resource "aws_s3_bucket_policy" "www" {
  bucket = local.bucket
  policy = jsonencode({
    Id      = "PolicyForCloudFrontPrivateContent"
    Version = "2008-10-17"
    Statement = [{
      Sid       = "AllowCloudFrontServicePrincipal"
      Effect    = "Allow"
      Principal = { Service = "cloudfront.amazonaws.com" }
      Action    = "s3:GetObject"
      Resource  = "arn:aws:s3:::${local.bucket}/*"
      Condition = {
        StringEquals = {
          "AWS:SourceArn" = aws_cloudfront_distribution.www.arn
        }
      }
    }]
  })
}

