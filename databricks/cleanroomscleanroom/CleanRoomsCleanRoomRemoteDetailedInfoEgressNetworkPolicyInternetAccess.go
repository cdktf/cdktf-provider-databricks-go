// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomscleanroom


type CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/clean_rooms_clean_room#allowed_internet_destinations CleanRoomsCleanRoom#allowed_internet_destinations}.
	AllowedInternetDestinations interface{} `field:"optional" json:"allowedInternetDestinations" yaml:"allowedInternetDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/clean_rooms_clean_room#allowed_storage_destinations CleanRoomsCleanRoom#allowed_storage_destinations}.
	AllowedStorageDestinations interface{} `field:"optional" json:"allowedStorageDestinations" yaml:"allowedStorageDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/clean_rooms_clean_room#log_only_mode CleanRoomsCleanRoom#log_only_mode}.
	LogOnlyMode *CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyMode `field:"optional" json:"logOnlyMode" yaml:"logOnlyMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/clean_rooms_clean_room#restriction_mode CleanRoomsCleanRoom#restriction_mode}.
	RestrictionMode *string `field:"optional" json:"restrictionMode" yaml:"restrictionMode"`
}

