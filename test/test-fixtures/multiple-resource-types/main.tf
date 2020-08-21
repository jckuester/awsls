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

resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
    awsls = "test-acc"
  }
}

resource "aws_subnet" "test" {
  vpc_id = aws_vpc.test.id
  cidr_block = "10.0.1.0/16"

  tags = {
    awsls = "test-acc"
  }
}