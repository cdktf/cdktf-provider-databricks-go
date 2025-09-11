// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomassetrevisionscleanroomasset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#asset_type DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#asset_type}.
	AssetType *string `field:"required" json:"assetType" yaml:"assetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#name DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#clean_room_name DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#clean_room_name}.
	CleanRoomName *string `field:"optional" json:"cleanRoomName" yaml:"cleanRoomName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#foreign_table DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#foreign_table}.
	ForeignTable *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetForeignTable `field:"optional" json:"foreignTable" yaml:"foreignTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#foreign_table_local_details DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#foreign_table_local_details}.
	ForeignTableLocalDetails *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetForeignTableLocalDetails `field:"optional" json:"foreignTableLocalDetails" yaml:"foreignTableLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#notebook DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#notebook}.
	Notebook *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetNotebook `field:"optional" json:"notebook" yaml:"notebook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#table DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#table}.
	Table *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetTable `field:"optional" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#table_local_details DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#table_local_details}.
	TableLocalDetails *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetTableLocalDetails `field:"optional" json:"tableLocalDetails" yaml:"tableLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#view DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#view}.
	View *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetView `field:"optional" json:"view" yaml:"view"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#view_local_details DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#view_local_details}.
	ViewLocalDetails *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetViewLocalDetails `field:"optional" json:"viewLocalDetails" yaml:"viewLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#volume_local_details DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#volume_local_details}.
	VolumeLocalDetails *DataDatabricksCleanRoomAssetRevisionsCleanRoomAssetVolumeLocalDetails `field:"optional" json:"volumeLocalDetails" yaml:"volumeLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset_revisions_clean_room_asset#workspace_id DataDatabricksCleanRoomAssetRevisionsCleanRoomAsset#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

