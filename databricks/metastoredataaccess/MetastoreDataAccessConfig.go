// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metastoredataaccess

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MetastoreDataAccessConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#name MetastoreDataAccess#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// aws_iam_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#aws_iam_role MetastoreDataAccess#aws_iam_role}
	AwsIamRole *MetastoreDataAccessAwsIamRole `field:"optional" json:"awsIamRole" yaml:"awsIamRole"`
	// azure_managed_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#azure_managed_identity MetastoreDataAccess#azure_managed_identity}
	AzureManagedIdentity *MetastoreDataAccessAzureManagedIdentity `field:"optional" json:"azureManagedIdentity" yaml:"azureManagedIdentity"`
	// azure_service_principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#azure_service_principal MetastoreDataAccess#azure_service_principal}
	AzureServicePrincipal *MetastoreDataAccessAzureServicePrincipal `field:"optional" json:"azureServicePrincipal" yaml:"azureServicePrincipal"`
	// cloudflare_api_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#cloudflare_api_token MetastoreDataAccess#cloudflare_api_token}
	CloudflareApiToken *MetastoreDataAccessCloudflareApiToken `field:"optional" json:"cloudflareApiToken" yaml:"cloudflareApiToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#comment MetastoreDataAccess#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// databricks_gcp_service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#databricks_gcp_service_account MetastoreDataAccess#databricks_gcp_service_account}
	DatabricksGcpServiceAccount *MetastoreDataAccessDatabricksGcpServiceAccount `field:"optional" json:"databricksGcpServiceAccount" yaml:"databricksGcpServiceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#force_destroy MetastoreDataAccess#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#force_update MetastoreDataAccess#force_update}.
	ForceUpdate interface{} `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// gcp_service_account_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#gcp_service_account_key MetastoreDataAccess#gcp_service_account_key}
	GcpServiceAccountKey *MetastoreDataAccessGcpServiceAccountKey `field:"optional" json:"gcpServiceAccountKey" yaml:"gcpServiceAccountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#id MetastoreDataAccess#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#is_default MetastoreDataAccess#is_default}.
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#isolation_mode MetastoreDataAccess#isolation_mode}.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#metastore_id MetastoreDataAccess#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#owner MetastoreDataAccess#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#read_only MetastoreDataAccess#read_only}.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/metastore_data_access#skip_validation MetastoreDataAccess#skip_validation}.
	SkipValidation interface{} `field:"optional" json:"skipValidation" yaml:"skipValidation"`
}

