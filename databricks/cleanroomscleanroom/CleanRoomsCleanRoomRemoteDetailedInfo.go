// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomscleanroom


type CleanRoomsCleanRoomRemoteDetailedInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/clean_rooms_clean_room#cloud_vendor CleanRoomsCleanRoom#cloud_vendor}.
	CloudVendor *string `field:"optional" json:"cloudVendor" yaml:"cloudVendor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/clean_rooms_clean_room#collaborators CleanRoomsCleanRoom#collaborators}.
	Collaborators interface{} `field:"optional" json:"collaborators" yaml:"collaborators"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/clean_rooms_clean_room#egress_network_policy CleanRoomsCleanRoom#egress_network_policy}.
	EgressNetworkPolicy *CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicy `field:"optional" json:"egressNetworkPolicy" yaml:"egressNetworkPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/clean_rooms_clean_room#region CleanRoomsCleanRoom#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

