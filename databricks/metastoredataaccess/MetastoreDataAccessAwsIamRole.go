// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metastoredataaccess


type MetastoreDataAccessAwsIamRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/metastore_data_access#role_arn MetastoreDataAccess#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/metastore_data_access#external_id MetastoreDataAccess#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/metastore_data_access#unity_catalog_iam_arn MetastoreDataAccess#unity_catalog_iam_arn}.
	UnityCatalogIamArn *string `field:"optional" json:"unityCatalogIamArn" yaml:"unityCatalogIamArn"`
}

