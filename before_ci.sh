#!/usr/bin/env bash

mkdir -p ~/.aws

# create AWS profile from AWS credential environment variables
cat > ~/.aws/credentials << EOL
[myaccount1]
aws_access_key_id = ${AWS_ACCESS_KEY_ID}
aws_secret_access_key = ${AWS_SECRET_ACCESS_KEY}

[myaccount2]
aws_access_key_id = ${AWS_ACCESS_KEY_ID2}
aws_secret_access_key = ${AWS_SECRET_ACCESS_KEY2}
EOL

cat > ~/.aws/config << EOL
[profile myaccount1]
region = us-west-2
output = json

[profile myaccount2]
region = us-east-1
output = json
EOL

# install Terraform
curl -sLo /tmp/terraform.zip https://releases.hashicorp.com/terraform/0.12.21/terraform_0.12.21_linux_amd64.zip
unzip /tmp/terraform.zip -d /tmp
mkdir -p ~/bin
mv /tmp/terraform ~/bin
export PATH="~/bin:$PATH"
