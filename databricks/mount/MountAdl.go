// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mount


type MountAdl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/mount#client_id Mount#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/mount#client_secret_key Mount#client_secret_key}.
	ClientSecretKey *string `field:"required" json:"clientSecretKey" yaml:"clientSecretKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/mount#client_secret_scope Mount#client_secret_scope}.
	ClientSecretScope *string `field:"required" json:"clientSecretScope" yaml:"clientSecretScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/mount#directory Mount#directory}.
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/mount#spark_conf_prefix Mount#spark_conf_prefix}.
	SparkConfPrefix *string `field:"optional" json:"sparkConfPrefix" yaml:"sparkConfPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/mount#storage_resource_name Mount#storage_resource_name}.
	StorageResourceName *string `field:"optional" json:"storageResourceName" yaml:"storageResourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/mount#tenant_id Mount#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

