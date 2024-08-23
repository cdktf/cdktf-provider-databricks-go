// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageCredentialConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#name StorageCredential#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// aws_iam_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#aws_iam_role StorageCredential#aws_iam_role}
	AwsIamRole *StorageCredentialAwsIamRole `field:"optional" json:"awsIamRole" yaml:"awsIamRole"`
	// azure_managed_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#azure_managed_identity StorageCredential#azure_managed_identity}
	AzureManagedIdentity *StorageCredentialAzureManagedIdentity `field:"optional" json:"azureManagedIdentity" yaml:"azureManagedIdentity"`
	// azure_service_principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#azure_service_principal StorageCredential#azure_service_principal}
	AzureServicePrincipal *StorageCredentialAzureServicePrincipal `field:"optional" json:"azureServicePrincipal" yaml:"azureServicePrincipal"`
	// cloudflare_api_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#cloudflare_api_token StorageCredential#cloudflare_api_token}
	CloudflareApiToken *StorageCredentialCloudflareApiToken `field:"optional" json:"cloudflareApiToken" yaml:"cloudflareApiToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#comment StorageCredential#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// databricks_gcp_service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#databricks_gcp_service_account StorageCredential#databricks_gcp_service_account}
	DatabricksGcpServiceAccount *StorageCredentialDatabricksGcpServiceAccount `field:"optional" json:"databricksGcpServiceAccount" yaml:"databricksGcpServiceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#force_destroy StorageCredential#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#force_update StorageCredential#force_update}.
	ForceUpdate interface{} `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// gcp_service_account_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#gcp_service_account_key StorageCredential#gcp_service_account_key}
	GcpServiceAccountKey *StorageCredentialGcpServiceAccountKey `field:"optional" json:"gcpServiceAccountKey" yaml:"gcpServiceAccountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#id StorageCredential#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#isolation_mode StorageCredential#isolation_mode}.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#metastore_id StorageCredential#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#owner StorageCredential#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#read_only StorageCredential#read_only}.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/storage_credential#skip_validation StorageCredential#skip_validation}.
	SkipValidation interface{} `field:"optional" json:"skipValidation" yaml:"skipValidation"`
}

