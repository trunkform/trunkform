variable "domain_name" {
  type = string
}

variable "environment" {
  type = string
}

resource "aws_route53_zone" "this" {
  name = var.domain_name
}

output "domain_name" {
  value = var.domain_name
}

output "environment" {
  value = var.environment
}

output "name_servers" {
  value = aws_route53_zone.this.name_servers
}

output "zone_id" {
  value = aws_route53_zone.this.zone_id
}
