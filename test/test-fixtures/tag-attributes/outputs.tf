output "vpc_id_single_tag" {
  value = aws_vpc.single_tag.id
}

output "vpc_id_multiple_tags" {
  value = aws_vpc.multiple_tags.id
}