terraform {
	backend "s3" {
		bucket="liferay-terraform-state-20250820205537886700000001"
		key="prod/dependencies/terraform.tfstate"
		region="us-east-2"
		use_lockfile=true
	}
}