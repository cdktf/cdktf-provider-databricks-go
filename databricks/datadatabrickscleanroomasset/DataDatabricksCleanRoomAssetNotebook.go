// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomasset


type DataDatabricksCleanRoomAssetNotebook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/data-sources/clean_room_asset#notebook_content DataDatabricksCleanRoomAsset#notebook_content}.
	NotebookContent *string `field:"required" json:"notebookContent" yaml:"notebookContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/data-sources/clean_room_asset#runner_collaborator_aliases DataDatabricksCleanRoomAsset#runner_collaborator_aliases}.
	RunnerCollaboratorAliases *[]*string `field:"optional" json:"runnerCollaboratorAliases" yaml:"runnerCollaboratorAliases"`
}

