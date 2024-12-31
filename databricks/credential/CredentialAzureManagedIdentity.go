// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package credential


type CredentialAzureManagedIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/credential#access_connector_id Credential#access_connector_id}.
	AccessConnectorId *string `field:"required" json:"accessConnectorId" yaml:"accessConnectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/credential#credential_id Credential#credential_id}.
	CredentialId *string `field:"optional" json:"credentialId" yaml:"credentialId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/credential#managed_identity_id Credential#managed_identity_id}.
	ManagedIdentityId *string `field:"optional" json:"managedIdentityId" yaml:"managedIdentityId"`
}

