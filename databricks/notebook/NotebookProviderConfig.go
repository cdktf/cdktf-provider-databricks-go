// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebook


type NotebookProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/notebook#workspace_id Notebook#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

