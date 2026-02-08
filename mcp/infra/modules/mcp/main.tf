terraform {
  backend "local" {}
}

variable "environment" {
  type = string
}

output "environment" {
  value = var.environment
}
