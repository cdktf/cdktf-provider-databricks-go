// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package azureadlsgen2mount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AzureAdlsGen2MountConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/azure_adls_gen2_mount#client_id AzureAdlsGen2Mount#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/azure_adls_gen2_mount#client_secret_key AzureAdlsGen2Mount#client_secret_key}.
	ClientSecretKey *string `field:"required" json:"clientSecretKey" yaml:"clientSecretKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/azure_adls_gen2_mount#client_secret_scope AzureAdlsGen2Mount#client_secret_scope}.
	ClientSecretScope *string `field:"required" json:"clientSecretScope" yaml:"clientSecretScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/azure_adls_gen2_mount#container_name AzureAdlsGen2Mount#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/azure_adls_gen2_mount#initialize_file_system AzureAdlsGen2Mount#initialize_file_system}.
	InitializeFileSystem interface{} `field:"required" json:"initializeFileSystem" yaml:"initializeFileSystem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/azure_adls_gen2_mount#mount_name AzureAdlsGen2Mount#mount_name}.
	MountName *string `field:"required" json:"mountName" yaml:"mountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/azure_adls_gen2_mount#storage_account_name AzureAdlsGen2Mount#storage_account_name}.
	StorageAccountName *string `field:"required" json:"storageAccountName" yaml:"storageAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/azure_adls_gen2_mount#tenant_id AzureAdlsGen2Mount#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/azure_adls_gen2_mount#cluster_id AzureAdlsGen2Mount#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/azure_adls_gen2_mount#directory AzureAdlsGen2Mount#directory}.
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/azure_adls_gen2_mount#id AzureAdlsGen2Mount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

