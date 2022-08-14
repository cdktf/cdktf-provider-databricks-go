// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type MountWasb struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#auth_type Mount#auth_type}.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#token_secret_key Mount#token_secret_key}.
	TokenSecretKey *string `field:"required" json:"tokenSecretKey" yaml:"tokenSecretKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#token_secret_scope Mount#token_secret_scope}.
	TokenSecretScope *string `field:"required" json:"tokenSecretScope" yaml:"tokenSecretScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#container_name Mount#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#directory Mount#directory}.
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#storage_account_name Mount#storage_account_name}.
	StorageAccountName *string `field:"optional" json:"storageAccountName" yaml:"storageAccountName"`
}

