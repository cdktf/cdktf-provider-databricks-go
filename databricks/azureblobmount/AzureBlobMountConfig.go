// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package azureblobmount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AzureBlobMountConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/azure_blob_mount#auth_type AzureBlobMount#auth_type}.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/azure_blob_mount#container_name AzureBlobMount#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/azure_blob_mount#mount_name AzureBlobMount#mount_name}.
	MountName *string `field:"required" json:"mountName" yaml:"mountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/azure_blob_mount#storage_account_name AzureBlobMount#storage_account_name}.
	StorageAccountName *string `field:"required" json:"storageAccountName" yaml:"storageAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/azure_blob_mount#token_secret_key AzureBlobMount#token_secret_key}.
	TokenSecretKey *string `field:"required" json:"tokenSecretKey" yaml:"tokenSecretKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/azure_blob_mount#token_secret_scope AzureBlobMount#token_secret_scope}.
	TokenSecretScope *string `field:"required" json:"tokenSecretScope" yaml:"tokenSecretScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/azure_blob_mount#cluster_id AzureBlobMount#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/azure_blob_mount#directory AzureBlobMount#directory}.
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/azure_blob_mount#id AzureBlobMount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

