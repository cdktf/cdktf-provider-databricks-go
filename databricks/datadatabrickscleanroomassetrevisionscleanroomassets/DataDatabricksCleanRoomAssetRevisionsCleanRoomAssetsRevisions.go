// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomassetrevisionscleanroomassets


type DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#asset_type DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#asset_type}.
	AssetType *string `field:"required" json:"assetType" yaml:"assetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#name DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#clean_room_name DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#clean_room_name}.
	CleanRoomName *string `field:"optional" json:"cleanRoomName" yaml:"cleanRoomName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#foreign_table DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#foreign_table}.
	ForeignTable *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTable `field:"optional" json:"foreignTable" yaml:"foreignTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#foreign_table_local_details DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#foreign_table_local_details}.
	ForeignTableLocalDetails *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsForeignTableLocalDetails `field:"optional" json:"foreignTableLocalDetails" yaml:"foreignTableLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#notebook DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#notebook}.
	Notebook *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsNotebook `field:"optional" json:"notebook" yaml:"notebook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#table DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#table}.
	Table *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTable `field:"optional" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#table_local_details DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#table_local_details}.
	TableLocalDetails *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsTableLocalDetails `field:"optional" json:"tableLocalDetails" yaml:"tableLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#view DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#view}.
	View *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsView `field:"optional" json:"view" yaml:"view"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#view_local_details DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#view_local_details}.
	ViewLocalDetails *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsViewLocalDetails `field:"optional" json:"viewLocalDetails" yaml:"viewLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_assets#volume_local_details DataDatabricksCleanRoomAssetRevisionsCleanRoomAssets#volume_local_details}.
	VolumeLocalDetails *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetsRevisionsVolumeLocalDetails `field:"optional" json:"volumeLocalDetails" yaml:"volumeLocalDetails"`
}

