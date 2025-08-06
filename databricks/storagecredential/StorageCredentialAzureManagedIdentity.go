// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecredential


type StorageCredentialAzureManagedIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/storage_credential#access_connector_id StorageCredential#access_connector_id}.
	AccessConnectorId *string `field:"required" json:"accessConnectorId" yaml:"accessConnectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/storage_credential#credential_id StorageCredential#credential_id}.
	CredentialId *string `field:"optional" json:"credentialId" yaml:"credentialId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/storage_credential#managed_identity_id StorageCredential#managed_identity_id}.
	ManagedIdentityId *string `field:"optional" json:"managedIdentityId" yaml:"managedIdentityId"`
}

