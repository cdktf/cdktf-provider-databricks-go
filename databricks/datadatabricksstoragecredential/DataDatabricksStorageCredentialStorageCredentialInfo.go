// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksstoragecredential


type DataDatabricksStorageCredentialStorageCredentialInfo struct {
	// aws_iam_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#aws_iam_role DataDatabricksStorageCredential#aws_iam_role}
	AwsIamRole *DataDatabricksStorageCredentialStorageCredentialInfoAwsIamRole `field:"optional" json:"awsIamRole" yaml:"awsIamRole"`
	// azure_managed_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#azure_managed_identity DataDatabricksStorageCredential#azure_managed_identity}
	AzureManagedIdentity *DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity `field:"optional" json:"azureManagedIdentity" yaml:"azureManagedIdentity"`
	// azure_service_principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#azure_service_principal DataDatabricksStorageCredential#azure_service_principal}
	AzureServicePrincipal *DataDatabricksStorageCredentialStorageCredentialInfoAzureServicePrincipal `field:"optional" json:"azureServicePrincipal" yaml:"azureServicePrincipal"`
	// cloudflare_api_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#cloudflare_api_token DataDatabricksStorageCredential#cloudflare_api_token}
	CloudflareApiToken *DataDatabricksStorageCredentialStorageCredentialInfoCloudflareApiToken `field:"optional" json:"cloudflareApiToken" yaml:"cloudflareApiToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#comment DataDatabricksStorageCredential#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#created_at DataDatabricksStorageCredential#created_at}.
	CreatedAt *float64 `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#created_by DataDatabricksStorageCredential#created_by}.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// databricks_gcp_service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#databricks_gcp_service_account DataDatabricksStorageCredential#databricks_gcp_service_account}
	DatabricksGcpServiceAccount *DataDatabricksStorageCredentialStorageCredentialInfoDatabricksGcpServiceAccount `field:"optional" json:"databricksGcpServiceAccount" yaml:"databricksGcpServiceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#id DataDatabricksStorageCredential#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#isolation_mode DataDatabricksStorageCredential#isolation_mode}.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#metastore_id DataDatabricksStorageCredential#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#name DataDatabricksStorageCredential#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#owner DataDatabricksStorageCredential#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#read_only DataDatabricksStorageCredential#read_only}.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#updated_at DataDatabricksStorageCredential#updated_at}.
	UpdatedAt *float64 `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#updated_by DataDatabricksStorageCredential#updated_by}.
	UpdatedBy *string `field:"optional" json:"updatedBy" yaml:"updatedBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/storage_credential#used_for_managed_storage DataDatabricksStorageCredential#used_for_managed_storage}.
	UsedForManagedStorage interface{} `field:"optional" json:"usedForManagedStorage" yaml:"usedForManagedStorage"`
}

