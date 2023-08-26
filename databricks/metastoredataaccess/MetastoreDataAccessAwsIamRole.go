// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metastoredataaccess


type MetastoreDataAccessAwsIamRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.0/docs/resources/metastore_data_access#role_arn MetastoreDataAccess#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

