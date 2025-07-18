// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomscleanroom


type CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/clean_rooms_clean_room#log_only_mode_type CleanRoomsCleanRoom#log_only_mode_type}.
	LogOnlyModeType *string `field:"optional" json:"logOnlyModeType" yaml:"logOnlyModeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/clean_rooms_clean_room#workloads CleanRoomsCleanRoom#workloads}.
	Workloads *[]*string `field:"optional" json:"workloads" yaml:"workloads"`
}

