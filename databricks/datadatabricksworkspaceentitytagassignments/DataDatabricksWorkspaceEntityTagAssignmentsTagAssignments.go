// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspaceentitytagassignments


type DataDatabricksWorkspaceEntityTagAssignmentsTagAssignments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/workspace_entity_tag_assignments#entity_id DataDatabricksWorkspaceEntityTagAssignments#entity_id}.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/workspace_entity_tag_assignments#entity_type DataDatabricksWorkspaceEntityTagAssignments#entity_type}.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/workspace_entity_tag_assignments#tag_key DataDatabricksWorkspaceEntityTagAssignments#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
}

