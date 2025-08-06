// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomassets


type DataDatabricksCleanRoomAssetsAssets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_assets#asset_type DataDatabricksCleanRoomAssets#asset_type}.
	AssetType *string `field:"required" json:"assetType" yaml:"assetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_assets#name DataDatabricksCleanRoomAssets#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_assets#clean_room_name DataDatabricksCleanRoomAssets#clean_room_name}.
	CleanRoomName *string `field:"optional" json:"cleanRoomName" yaml:"cleanRoomName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_assets#foreign_table DataDatabricksCleanRoomAssets#foreign_table}.
	ForeignTable *DataDatabricksCleanRoomAssetsAssetsForeignTable `field:"optional" json:"foreignTable" yaml:"foreignTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_assets#foreign_table_local_details DataDatabricksCleanRoomAssets#foreign_table_local_details}.
	ForeignTableLocalDetails *DataDatabricksCleanRoomAssetsAssetsForeignTableLocalDetails `field:"optional" json:"foreignTableLocalDetails" yaml:"foreignTableLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_assets#notebook DataDatabricksCleanRoomAssets#notebook}.
	Notebook *DataDatabricksCleanRoomAssetsAssetsNotebook `field:"optional" json:"notebook" yaml:"notebook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_assets#table DataDatabricksCleanRoomAssets#table}.
	Table *DataDatabricksCleanRoomAssetsAssetsTable `field:"optional" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_assets#table_local_details DataDatabricksCleanRoomAssets#table_local_details}.
	TableLocalDetails *DataDatabricksCleanRoomAssetsAssetsTableLocalDetails `field:"optional" json:"tableLocalDetails" yaml:"tableLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_assets#view DataDatabricksCleanRoomAssets#view}.
	View *DataDatabricksCleanRoomAssetsAssetsView `field:"optional" json:"view" yaml:"view"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_assets#view_local_details DataDatabricksCleanRoomAssets#view_local_details}.
	ViewLocalDetails *DataDatabricksCleanRoomAssetsAssetsViewLocalDetails `field:"optional" json:"viewLocalDetails" yaml:"viewLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_assets#volume_local_details DataDatabricksCleanRoomAssets#volume_local_details}.
	VolumeLocalDetails *DataDatabricksCleanRoomAssetsAssetsVolumeLocalDetails `field:"optional" json:"volumeLocalDetails" yaml:"volumeLocalDetails"`
}

