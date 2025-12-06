// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package permissionassignment


type PermissionAssignmentProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/permission_assignment#workspace_id PermissionAssignment#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

