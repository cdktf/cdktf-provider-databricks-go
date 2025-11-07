// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksstoragecredential


type DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/storage_credential#access_connector_id DataDatabricksStorageCredential#access_connector_id}.
	AccessConnectorId *string `field:"required" json:"accessConnectorId" yaml:"accessConnectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/storage_credential#credential_id DataDatabricksStorageCredential#credential_id}.
	CredentialId *string `field:"optional" json:"credentialId" yaml:"credentialId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/storage_credential#managed_identity_id DataDatabricksStorageCredential#managed_identity_id}.
	ManagedIdentityId *string `field:"optional" json:"managedIdentityId" yaml:"managedIdentityId"`
}

