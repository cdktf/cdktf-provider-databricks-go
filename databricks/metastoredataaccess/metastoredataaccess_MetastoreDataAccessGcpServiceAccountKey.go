package metastoredataaccess


type MetastoreDataAccessGcpServiceAccountKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#email MetastoreDataAccess#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#private_key MetastoreDataAccess#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#private_key_id MetastoreDataAccess#private_key_id}.
	PrivateKeyId *string `field:"required" json:"privateKeyId" yaml:"privateKeyId"`
}

