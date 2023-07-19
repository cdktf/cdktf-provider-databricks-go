package metastoredataaccess


type MetastoreDataAccessAzureManagedIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.21.0/docs/resources/metastore_data_access#access_connector_id MetastoreDataAccess#access_connector_id}.
	AccessConnectorId *string `field:"required" json:"accessConnectorId" yaml:"accessConnectorId"`
}

