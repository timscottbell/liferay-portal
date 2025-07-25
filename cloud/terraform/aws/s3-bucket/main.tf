module "s3_bucket" {
  source = "../modules/s3-bucket"

  deployment_name = var.deployment_name
}