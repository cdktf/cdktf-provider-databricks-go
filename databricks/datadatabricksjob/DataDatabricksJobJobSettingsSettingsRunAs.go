package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsRunAs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/data-sources/job#service_principal_name DataDatabricksJob#service_principal_name}.
	ServicePrincipalName *string `field:"optional" json:"servicePrincipalName" yaml:"servicePrincipalName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/data-sources/job#user_name DataDatabricksJob#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

