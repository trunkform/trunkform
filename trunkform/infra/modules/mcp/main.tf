variable "environment" {
  type = string
}

output "environment" {
  value = var.environment
}

# resource "aws_s3_bucket" "this" {
#   bucket = "mcp-trunkform-${var.environment}"
# }
