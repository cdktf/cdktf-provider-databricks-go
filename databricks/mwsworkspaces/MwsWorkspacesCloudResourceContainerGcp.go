// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsworkspaces


type MwsWorkspacesCloudResourceContainerGcp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.1/docs/resources/mws_workspaces#project_id MwsWorkspaces#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

