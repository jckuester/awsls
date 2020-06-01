#!/usr/bin/env bash

mkdir -p ~/.aws

# create AWS profile from AWS credential environment variables
cat > ~/.aws/credentials << EOL
[awsls]
aws_access_key_id = ${AWS_ACCESS_KEY_ID}
aws_secret_access_key = ${AWS_SECRET_ACCESS_KEY}
EOL

# install Terraform
curl -sLo /tmp/terraform.zip https://releases.hashicorp.com/terraform/0.12.21/terraform_0.12.21_linux_amd64.zip
unzip /tmp/terraform.zip -d /tmp
mkdir -p ~/bin
mv /tmp/terraform ~/bin
export PATH="~/bin:$PATH"
