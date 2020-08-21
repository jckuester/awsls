output "vpc_id" {
  value = aws_vpc.test.id
}

output "subnet_id" {
  value = aws_subnet.test.id
}