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
    name = "dns1.registrar-servers.com"
  }
  name_server {
    name = "dns2.registrar-servers.com"
  }
}
