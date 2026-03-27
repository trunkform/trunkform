variable "domain_name"           { type = string }
variable "zone_id"                { type = string }
variable "cloudfront_domain_name" { type = string }
variable "cloudfront_zone_id"     { type = string }

resource "aws_route53_record" "a" {
  name    = var.domain_name
  type    = "A"
  zone_id = var.zone_id

  alias {
    name                   = var.cloudfront_domain_name
    zone_id                = var.cloudfront_zone_id
    evaluate_target_health = false
  }
}
