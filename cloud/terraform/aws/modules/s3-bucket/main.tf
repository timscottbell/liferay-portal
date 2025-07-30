locals {
  bucket_prefix = "${var.deployment_name}-s3-bucket-"
}

module "s3_bucket" {
  source  = "terraform-aws-modules/s3-bucket/aws"
  version = "~> 4.1.1"

  bucket = "${local.bucket_prefix}${random_password.s3_bucket_suffix.result}"

  block_public_acls   = true
  block_public_policy = true
  control_object_ownership = true
  force_destroy       = true
  ignore_public_acls  = true
  restrict_public_buckets = true
  object_ownership = "BucketOwnerPreferred"

  server_side_encryption_configuration = {
    rule = {
      apply_server_side_encryption_by_default = {
        sse_algorithm = "aws:kms"
      }
      bucket_key_enabled = true
    }
  }

  versioning = {
    enabled = true
  }

  tags = {
    Backup = "true"
  }
}

resource "random_password" "s3_bucket_suffix" {
  length  = 8
  special = false
  upper   = false
}