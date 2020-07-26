provider "aws" {
  version = "~> 2.0"

  profile = var.profile
  region  = var.region
}

terraform {
  # The configuration for this backend will be filled in by Terragrunt
  backend "s3" {
  }
}

resource "aws_vpc" "single_tag" {
  cidr_block = "10.0.0.0/16"

  tags = {
    foo = "bar"
  }
}

resource "aws_vpc" "multiple_tags" {
  cidr_block = "10.0.0.0/16"

  tags = {
    foo = "bar"
    bar = "baz"
  }
}