// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomasset


type DataDatabricksCleanRoomAssetForeignTableColumnsMask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#function_name DataDatabricksCleanRoomAsset#function_name}.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/clean_room_asset#using_column_names DataDatabricksCleanRoomAsset#using_column_names}.
	UsingColumnNames *[]*string `field:"optional" json:"usingColumnNames" yaml:"usingColumnNames"`
}

