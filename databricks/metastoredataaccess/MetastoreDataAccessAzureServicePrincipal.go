package metastoredataaccess


type MetastoreDataAccessAzureServicePrincipal struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.0/docs/resources/metastore_data_access#application_id MetastoreDataAccess#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.0/docs/resources/metastore_data_access#client_secret MetastoreDataAccess#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.0/docs/resources/metastore_data_access#directory_id MetastoreDataAccess#directory_id}.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
}

