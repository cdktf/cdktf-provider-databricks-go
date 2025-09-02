// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomassetrevisionscleanroomasset


type DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetNotebook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#notebook_content DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#notebook_content}.
	NotebookContent *string `field:"required" json:"notebookContent" yaml:"notebookContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#runner_collaborator_aliases DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#runner_collaborator_aliases}.
	RunnerCollaboratorAliases *[]*string `field:"optional" json:"runnerCollaboratorAliases" yaml:"runnerCollaboratorAliases"`
}

