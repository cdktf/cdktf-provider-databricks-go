// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksstoragecredential


type DataDatabricksStorageCredentialStorageCredentialInfoAwsIamRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/storage_credential#role_arn DataDatabricksStorageCredential#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/storage_credential#external_id DataDatabricksStorageCredential#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/data-sources/storage_credential#unity_catalog_iam_arn DataDatabricksStorageCredential#unity_catalog_iam_arn}.
	UnityCatalogIamArn *string `field:"optional" json:"unityCatalogIamArn" yaml:"unityCatalogIamArn"`
}

