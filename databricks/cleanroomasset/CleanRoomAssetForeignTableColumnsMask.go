// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomasset


type CleanRoomAssetForeignTableColumnsMask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#function_name CleanRoomAsset#function_name}.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/clean_room_asset#using_column_names CleanRoomAsset#using_column_names}.
	UsingColumnNames *[]*string `field:"optional" json:"usingColumnNames" yaml:"usingColumnNames"`
}

