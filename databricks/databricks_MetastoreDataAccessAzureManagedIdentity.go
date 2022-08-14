// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type MetastoreDataAccessAzureManagedIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#access_connector_id MetastoreDataAccess#access_connector_id}.
	AccessConnectorId *string `field:"required" json:"accessConnectorId" yaml:"accessConnectorId"`
}

