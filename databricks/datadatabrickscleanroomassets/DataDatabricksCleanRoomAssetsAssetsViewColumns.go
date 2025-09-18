// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomassets


type DataDatabricksCleanRoomAssetsAssetsViewColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#comment DataDatabricksCleanRoomAssets#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#mask DataDatabricksCleanRoomAssets#mask}.
	Mask *DataDatabricksCleanRoomAssetsAssetsViewColumnsMask `field:"optional" json:"mask" yaml:"mask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#name DataDatabricksCleanRoomAssets#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#nullable DataDatabricksCleanRoomAssets#nullable}.
	Nullable interface{} `field:"optional" json:"nullable" yaml:"nullable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#partition_index DataDatabricksCleanRoomAssets#partition_index}.
	PartitionIndex *float64 `field:"optional" json:"partitionIndex" yaml:"partitionIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#position DataDatabricksCleanRoomAssets#position}.
	Position *float64 `field:"optional" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#type_interval_type DataDatabricksCleanRoomAssets#type_interval_type}.
	TypeIntervalType *string `field:"optional" json:"typeIntervalType" yaml:"typeIntervalType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#type_json DataDatabricksCleanRoomAssets#type_json}.
	TypeJson *string `field:"optional" json:"typeJson" yaml:"typeJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#type_name DataDatabricksCleanRoomAssets#type_name}.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#type_precision DataDatabricksCleanRoomAssets#type_precision}.
	TypePrecision *float64 `field:"optional" json:"typePrecision" yaml:"typePrecision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#type_scale DataDatabricksCleanRoomAssets#type_scale}.
	TypeScale *float64 `field:"optional" json:"typeScale" yaml:"typeScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_room_assets#type_text DataDatabricksCleanRoomAssets#type_text}.
	TypeText *string `field:"optional" json:"typeText" yaml:"typeText"`
}

