// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type MountAbfs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#client_id Mount#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#client_secret_key Mount#client_secret_key}.
	ClientSecretKey *string `field:"required" json:"clientSecretKey" yaml:"clientSecretKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#client_secret_scope Mount#client_secret_scope}.
	ClientSecretScope *string `field:"required" json:"clientSecretScope" yaml:"clientSecretScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#initialize_file_system Mount#initialize_file_system}.
	InitializeFileSystem interface{} `field:"required" json:"initializeFileSystem" yaml:"initializeFileSystem"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#container_name Mount#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#directory Mount#directory}.
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#storage_account_name Mount#storage_account_name}.
	StorageAccountName *string `field:"optional" json:"storageAccountName" yaml:"storageAccountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#tenant_id Mount#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}
