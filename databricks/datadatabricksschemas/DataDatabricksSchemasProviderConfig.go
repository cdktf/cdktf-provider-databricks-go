// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksschemas


type DataDatabricksSchemasProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/schemas#workspace_id DataDatabricksSchemas#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

