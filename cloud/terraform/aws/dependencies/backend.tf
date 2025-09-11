terraform {
	backend "s3" {
		bucket="liferay-terraform-state-20250820205537886700000001"
		dynamodb_table="liferay-terraform-state-locks"
		key="prod/dependencies/terraform.tfstate"
		region="us-east-2"
	}
}