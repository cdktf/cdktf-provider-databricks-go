// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomasset


type DataDatabricksCleanRoomAssetTableLocalDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset#local_name DataDatabricksCleanRoomAsset#local_name}.
	LocalName *string `field:"required" json:"localName" yaml:"localName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_room_asset#partitions DataDatabricksCleanRoomAsset#partitions}.
	Partitions interface{} `field:"optional" json:"partitions" yaml:"partitions"`
}

