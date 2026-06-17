variable "environment" {
  type = string
}

variable "domain_name" {
  type = string
}

variable "name_servers" {
  type = list(string)
  default = [null, null, null, null]
}

removed {
  from = aws_route53_zone.this

  lifecycle {
    destroy = false
  }
}

resource "aws_route53domains_registered_domain" "this" {
  domain_name = var.domain_name
  auto_renew  = true

  admin_privacy      = true
  registrant_privacy = true
  tech_privacy       = true
  billing_privacy    = true

  name_server {
    name = var.name_servers[0]
  }
  name_server {
    name = var.name_servers[1]
  }
  name_server {
    name = var.name_servers[2]
  }
  name_server {
    name = var.name_servers[3]
  }
}

output "domain_name" {
  value = var.domain_name
}

output "environment" {
  value = var.environment
}
