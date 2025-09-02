// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecredential


type StorageCredentialDatabricksGcpServiceAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/storage_credential#credential_id StorageCredential#credential_id}.
	CredentialId *string `field:"optional" json:"credentialId" yaml:"credentialId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/storage_credential#email StorageCredential#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
}

