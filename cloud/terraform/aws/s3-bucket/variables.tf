variable "deployment_name" {
  description = "The base name for the deployment, used to prefix resource names."
  type        = string
}
variable "region" {
  description = "The AWS region."
  type        = string
}