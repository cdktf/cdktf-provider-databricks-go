package storagecredential


type StorageCredentialGcpServiceAccountKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/storage_credential#email StorageCredential#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/storage_credential#private_key StorageCredential#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/storage_credential#private_key_id StorageCredential#private_key_id}.
	PrivateKeyId *string `field:"required" json:"privateKeyId" yaml:"privateKeyId"`
}
