output "s3_bucket_id" {
  description = "The name of the S3 bucket."
  value       = module.s3_bucket.s3_bucket_id
}

output "s3_bucket_prefix" {
  description = "The prefix of the S3 bucket."
  value       = local.bucket_prefix
}

output "s3_bucket_region" {
  description = "The region of the S3 bucket."
  value       = module.s3_bucket.s3_bucket_region
}
