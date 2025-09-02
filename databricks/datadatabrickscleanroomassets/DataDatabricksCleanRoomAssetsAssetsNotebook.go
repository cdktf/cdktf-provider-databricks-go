// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomassets


type DataDatabricksCleanRoomAssetsAssetsNotebook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_assets#notebook_content DataDatabricksCleanRoomAssets#notebook_content}.
	NotebookContent *string `field:"required" json:"notebookContent" yaml:"notebookContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_assets#runner_collaborator_aliases DataDatabricksCleanRoomAssets#runner_collaborator_aliases}.
	RunnerCollaboratorAliases *[]*string `field:"optional" json:"runnerCollaboratorAliases" yaml:"runnerCollaboratorAliases"`
}

