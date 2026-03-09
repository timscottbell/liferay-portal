locals {
	common_labels={
		"app.kubernetes.io/managed-by"=local.terraform_manager_name
		"environment"="internal"
	}
	goldilocks_labels=merge(
		local.common_labels,
		{
			"goldilocks.fairwinds.com/enabled"="true"
		})
	terraform_manager_name="liferay-cloud-native-terraform"
}