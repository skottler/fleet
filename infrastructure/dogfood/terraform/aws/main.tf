variable "region" {
  default = "us-east-2"
}

provider "aws" {
  region = var.region
}


provider "tls" {
  # Configuration options
}


terraform {
  // these values should match what is bootstrapped in ./remote-state
  backend "s3" {
    bucket         = "fleet-terraform-remote-state"
    region         = "us-east-2"
    key            = "fleet"
    dynamodb_table = "fleet-terraform-state-lock"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.32.0"
    }

    tls = {
      source  = "hashicorp/tls"
      version = "3.3.0"
    }
  }
}

data "aws_caller_identity" "current" {}
