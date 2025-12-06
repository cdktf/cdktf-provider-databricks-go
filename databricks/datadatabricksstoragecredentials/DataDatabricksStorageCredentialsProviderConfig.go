// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksstoragecredentials


type DataDatabricksStorageCredentialsProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/storage_credentials#workspace_id DataDatabricksStorageCredentials#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

