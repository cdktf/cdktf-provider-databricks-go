// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomasset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCleanRoomAssetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#asset_type DataDatabricksCleanRoomAsset#asset_type}.
	AssetType *string `field:"required" json:"assetType" yaml:"assetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#name DataDatabricksCleanRoomAsset#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#clean_room_name DataDatabricksCleanRoomAsset#clean_room_name}.
	CleanRoomName *string `field:"optional" json:"cleanRoomName" yaml:"cleanRoomName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#foreign_table DataDatabricksCleanRoomAsset#foreign_table}.
	ForeignTable *DataDatabricksCleanRoomAssetForeignTable `field:"optional" json:"foreignTable" yaml:"foreignTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#foreign_table_local_details DataDatabricksCleanRoomAsset#foreign_table_local_details}.
	ForeignTableLocalDetails *DataDatabricksCleanRoomAssetForeignTableLocalDetails `field:"optional" json:"foreignTableLocalDetails" yaml:"foreignTableLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#notebook DataDatabricksCleanRoomAsset#notebook}.
	Notebook *DataDatabricksCleanRoomAssetNotebook `field:"optional" json:"notebook" yaml:"notebook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#table DataDatabricksCleanRoomAsset#table}.
	Table *DataDatabricksCleanRoomAssetTable `field:"optional" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#table_local_details DataDatabricksCleanRoomAsset#table_local_details}.
	TableLocalDetails *DataDatabricksCleanRoomAssetTableLocalDetails `field:"optional" json:"tableLocalDetails" yaml:"tableLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#view DataDatabricksCleanRoomAsset#view}.
	View *DataDatabricksCleanRoomAssetView `field:"optional" json:"view" yaml:"view"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#view_local_details DataDatabricksCleanRoomAsset#view_local_details}.
	ViewLocalDetails *DataDatabricksCleanRoomAssetViewLocalDetails `field:"optional" json:"viewLocalDetails" yaml:"viewLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#volume_local_details DataDatabricksCleanRoomAsset#volume_local_details}.
	VolumeLocalDetails *DataDatabricksCleanRoomAssetVolumeLocalDetails `field:"optional" json:"volumeLocalDetails" yaml:"volumeLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#workspace_id DataDatabricksCleanRoomAsset#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

