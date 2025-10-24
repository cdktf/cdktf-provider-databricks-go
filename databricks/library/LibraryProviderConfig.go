// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package library


type LibraryProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/library#workspace_id Library#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

