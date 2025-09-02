// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomasset


type CleanRoomAssetForeignTableColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#comment CleanRoomAsset#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#mask CleanRoomAsset#mask}.
	Mask *CleanRoomAssetForeignTableColumnsMask `field:"optional" json:"mask" yaml:"mask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#name CleanRoomAsset#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#nullable CleanRoomAsset#nullable}.
	Nullable interface{} `field:"optional" json:"nullable" yaml:"nullable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#partition_index CleanRoomAsset#partition_index}.
	PartitionIndex *float64 `field:"optional" json:"partitionIndex" yaml:"partitionIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#position CleanRoomAsset#position}.
	Position *float64 `field:"optional" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#type_interval_type CleanRoomAsset#type_interval_type}.
	TypeIntervalType *string `field:"optional" json:"typeIntervalType" yaml:"typeIntervalType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#type_json CleanRoomAsset#type_json}.
	TypeJson *string `field:"optional" json:"typeJson" yaml:"typeJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#type_name CleanRoomAsset#type_name}.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#type_precision CleanRoomAsset#type_precision}.
	TypePrecision *float64 `field:"optional" json:"typePrecision" yaml:"typePrecision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#type_scale CleanRoomAsset#type_scale}.
	TypeScale *float64 `field:"optional" json:"typeScale" yaml:"typeScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#type_text CleanRoomAsset#type_text}.
	TypeText *string `field:"optional" json:"typeText" yaml:"typeText"`
}

