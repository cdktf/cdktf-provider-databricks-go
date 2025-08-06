// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomasset


type CleanRoomAssetTableLocalDetailsPartitionsValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#name CleanRoomAsset#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#op CleanRoomAsset#op}.
	Op *string `field:"optional" json:"op" yaml:"op"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#recipient_property_key CleanRoomAsset#recipient_property_key}.
	RecipientPropertyKey *string `field:"optional" json:"recipientPropertyKey" yaml:"recipientPropertyKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#value CleanRoomAsset#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

