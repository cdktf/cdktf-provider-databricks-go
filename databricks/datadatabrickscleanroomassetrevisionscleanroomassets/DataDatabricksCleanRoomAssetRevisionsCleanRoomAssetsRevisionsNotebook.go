// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomassetrevisionscleanroomassets


type DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsNotebook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#notebook_content DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#notebook_content}.
	NotebookContent *string `field:"required" json:"notebookContent" yaml:"notebookContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#runner_collaborator_aliases DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#runner_collaborator_aliases}.
	RunnerCollaboratorAliases *[]*string `field:"optional" json:"runnerCollaboratorAliases" yaml:"runnerCollaboratorAliases"`
}

