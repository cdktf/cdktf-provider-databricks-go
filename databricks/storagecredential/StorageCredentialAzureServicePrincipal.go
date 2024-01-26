// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecredential


type StorageCredentialAzureServicePrincipal struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.35.0/docs/resources/storage_credential#application_id StorageCredential#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.35.0/docs/resources/storage_credential#client_secret StorageCredential#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.35.0/docs/resources/storage_credential#directory_id StorageCredential#directory_id}.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
}

