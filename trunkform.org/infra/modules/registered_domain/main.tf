variable "environment" {
  type = string
}

variable "domain_name" {
  type = string
}

output "environment" {
  value = var.environment
}

output "domain_name" {
  value = var.domain_name
}

output "zone_id" {
  value = aws_route53_zone.this.zone_id
}

resource "aws_route53_zone" "this" {
  name = var.domain_name
}

resource "aws_route53domains_registered_domain" "this" {
  domain_name = var.domain_name
  auto_renew  = true

  admin_privacy      = true
  registrant_privacy = true
  tech_privacy       = true
  billing_privacy    = true

  name_server {
    name = aws_route53_zone.this.name_servers[0]
  }
  name_server {
    name = aws_route53_zone.this.name_servers[1]
  }
  name_server {
    name = aws_route53_zone.this.name_servers[2]
  }
  name_server {
    name = aws_route53_zone.this.name_servers[3]
  }
}
