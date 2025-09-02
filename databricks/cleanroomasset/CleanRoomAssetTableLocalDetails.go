// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomasset


type CleanRoomAssetTableLocalDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#local_name CleanRoomAsset#local_name}.
	LocalName *string `field:"required" json:"localName" yaml:"localName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#partitions CleanRoomAsset#partitions}.
	Partitions interface{} `field:"optional" json:"partitions" yaml:"partitions"`
}

