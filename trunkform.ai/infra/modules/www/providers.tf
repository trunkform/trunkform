terraform {
  required_providers {
    aws = {
      source                = "hashicorp/aws"
      configuration_aliases = [aws.global]
    }
  }
}

provider "aws" {
  region = "us-east-2"
}

provider "aws" {
  alias  = "global"
  region = "us-east-1"
}
