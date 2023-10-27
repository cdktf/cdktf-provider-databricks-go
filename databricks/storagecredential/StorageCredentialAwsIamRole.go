// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecredential


type StorageCredentialAwsIamRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/resources/storage_credential#role_arn StorageCredential#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

