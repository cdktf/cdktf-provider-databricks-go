// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecredential


type StorageCredentialAzureManagedIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.0/docs/resources/storage_credential#access_connector_id StorageCredential#access_connector_id}.
	AccessConnectorId *string `field:"required" json:"accessConnectorId" yaml:"accessConnectorId"`
}

