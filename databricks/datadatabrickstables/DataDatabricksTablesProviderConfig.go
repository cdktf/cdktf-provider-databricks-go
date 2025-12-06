// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstables


type DataDatabricksTablesProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/tables#workspace_id DataDatabricksTables#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

