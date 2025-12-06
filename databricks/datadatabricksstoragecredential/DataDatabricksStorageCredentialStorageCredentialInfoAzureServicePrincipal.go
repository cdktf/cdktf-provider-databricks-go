// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksstoragecredential


type DataDatabricksStorageCredentialStorageCredentialInfoAzureServicePrincipal struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/storage_credential#application_id DataDatabricksStorageCredential#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/storage_credential#client_secret DataDatabricksStorageCredential#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/storage_credential#directory_id DataDatabricksStorageCredential#directory_id}.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
}

