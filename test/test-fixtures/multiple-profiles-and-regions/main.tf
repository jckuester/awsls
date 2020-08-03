provider "aws" {
  version = "~> 2.0"
  alias  = "profile1-region1"

  profile = var.profile1
  region  = var.region1
}

provider "aws" {
  version = "~> 2.0"
  alias  = "profile1-region2"

  profile = var.profile1
  region  = var.region2
}

provider "aws" {
  version = "~> 2.0"
  alias  = "profile2-region1"

  profile = var.profile2
  region  = var.region1
}

provider "aws" {
  version = "~> 2.0"
  alias  = "profile2-region2"

  profile = var.profile2
  region  = var.region2
}

terraform {
  # The configuration for this backend will be filled in by Terragrunt
  backend "s3" {
  }
}

resource "aws_vpc" "test1" {
  provider = aws.profile1-region1

  cidr_block = "10.0.0.0/16"

  tags = {
    awsls = "test-acc"
    foo = "${var.profile1}-${var.region1}"
  }
}

resource "aws_vpc" "test2" {
  provider = aws.profile1-region2

  cidr_block = "10.0.0.0/16"

  tags = {
    awsls = "test-acc"
    foo = "${var.profile1}-${var.region2}"
  }
}

resource "aws_vpc" "test3" {
  provider = aws.profile2-region1

  cidr_block = "10.0.0.0/16"

  tags = {
    awsls = "test-acc"
    foo = "${var.profile2}-${var.region1}"
  }
}

resource "aws_vpc" "test4" {
  provider = aws.profile2-region2

  cidr_block = "10.0.0.0/16"

  tags = {
    awsls = "test-acc"
    foo = "${var.profile2}-${var.region2}"
  }
}