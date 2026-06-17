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

output "names_servers" {
  value = aws_route53_zone.this.name_servers
}

resource "aws_route53_zone" "this" {
  name = var.domain_name
}

import {
  to = aws_route53_zone.this
  id = "Z0362682YDFN6HWXCEER"
}
