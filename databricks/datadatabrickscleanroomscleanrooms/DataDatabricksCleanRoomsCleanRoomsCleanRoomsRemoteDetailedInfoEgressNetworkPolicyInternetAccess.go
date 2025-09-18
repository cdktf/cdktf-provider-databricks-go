// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomscleanrooms


type DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_rooms_clean_rooms#allowed_internet_destinations DataDatabricksCleanRoomsCleanRooms#allowed_internet_destinations}.
	AllowedInternetDestinations interface{} `field:"optional" json:"allowedInternetDestinations" yaml:"allowedInternetDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_rooms_clean_rooms#allowed_storage_destinations DataDatabricksCleanRoomsCleanRooms#allowed_storage_destinations}.
	AllowedStorageDestinations interface{} `field:"optional" json:"allowedStorageDestinations" yaml:"allowedStorageDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_rooms_clean_rooms#log_only_mode DataDatabricksCleanRoomsCleanRooms#log_only_mode}.
	LogOnlyMode *DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyMode `field:"optional" json:"logOnlyMode" yaml:"logOnlyMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/clean_rooms_clean_rooms#restriction_mode DataDatabricksCleanRoomsCleanRooms#restriction_mode}.
	RestrictionMode *string `field:"optional" json:"restrictionMode" yaml:"restrictionMode"`
}

