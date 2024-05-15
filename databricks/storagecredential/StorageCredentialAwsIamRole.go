// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecredential


type StorageCredentialAwsIamRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.43.0/docs/resources/storage_credential#role_arn StorageCredential#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.43.0/docs/resources/storage_credential#external_id StorageCredential#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.43.0/docs/resources/storage_credential#unity_catalog_iam_arn StorageCredential#unity_catalog_iam_arn}.
	UnityCatalogIamArn *string `field:"optional" json:"unityCatalogIamArn" yaml:"unityCatalogIamArn"`
}

