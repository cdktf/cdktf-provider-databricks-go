// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type MountAdl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#client_id Mount#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#client_secret_key Mount#client_secret_key}.
	ClientSecretKey *string `field:"required" json:"clientSecretKey" yaml:"clientSecretKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#client_secret_scope Mount#client_secret_scope}.
	ClientSecretScope *string `field:"required" json:"clientSecretScope" yaml:"clientSecretScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#directory Mount#directory}.
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#spark_conf_prefix Mount#spark_conf_prefix}.
	SparkConfPrefix *string `field:"optional" json:"sparkConfPrefix" yaml:"sparkConfPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#storage_resource_name Mount#storage_resource_name}.
	StorageResourceName *string `field:"optional" json:"storageResourceName" yaml:"storageResourceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#tenant_id Mount#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

