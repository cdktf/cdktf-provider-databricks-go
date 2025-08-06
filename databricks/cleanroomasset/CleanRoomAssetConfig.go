// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomasset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanRoomAssetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#asset_type CleanRoomAsset#asset_type}.
	AssetType *string `field:"required" json:"assetType" yaml:"assetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#name CleanRoomAsset#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#clean_room_name CleanRoomAsset#clean_room_name}.
	CleanRoomName *string `field:"optional" json:"cleanRoomName" yaml:"cleanRoomName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#foreign_table CleanRoomAsset#foreign_table}.
	ForeignTable *CleanRoomAssetForeignTable `field:"optional" json:"foreignTable" yaml:"foreignTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#foreign_table_local_details CleanRoomAsset#foreign_table_local_details}.
	ForeignTableLocalDetails *CleanRoomAssetForeignTableLocalDetails `field:"optional" json:"foreignTableLocalDetails" yaml:"foreignTableLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#notebook CleanRoomAsset#notebook}.
	Notebook *CleanRoomAssetNotebook `field:"optional" json:"notebook" yaml:"notebook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#table CleanRoomAsset#table}.
	Table *CleanRoomAssetTable `field:"optional" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#table_local_details CleanRoomAsset#table_local_details}.
	TableLocalDetails *CleanRoomAssetTableLocalDetails `field:"optional" json:"tableLocalDetails" yaml:"tableLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#view CleanRoomAsset#view}.
	View *CleanRoomAssetView `field:"optional" json:"view" yaml:"view"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#view_local_details CleanRoomAsset#view_local_details}.
	ViewLocalDetails *CleanRoomAssetViewLocalDetails `field:"optional" json:"viewLocalDetails" yaml:"viewLocalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#volume_local_details CleanRoomAsset#volume_local_details}.
	VolumeLocalDetails *CleanRoomAssetVolumeLocalDetails `field:"optional" json:"volumeLocalDetails" yaml:"volumeLocalDetails"`
}

