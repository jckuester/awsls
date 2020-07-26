output "vpc_id1" {
  value = aws_vpc.single_tag.id
}

output "vpc_id2" {
  value = aws_vpc.multiple_tags.id
}