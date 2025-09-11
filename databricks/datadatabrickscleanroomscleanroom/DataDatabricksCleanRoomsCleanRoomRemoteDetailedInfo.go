// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomscleanroom


type DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_room#cloud_vendor DataDatabricksCleanRoomsCleanRoom#cloud_vendor}.
	CloudVendor *string `field:"optional" json:"cloudVendor" yaml:"cloudVendor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_room#collaborators DataDatabricksCleanRoomsCleanRoom#collaborators}.
	Collaborators interface{} `field:"optional" json:"collaborators" yaml:"collaborators"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_room#egress_network_policy DataDatabricksCleanRoomsCleanRoom#egress_network_policy}.
	EgressNetworkPolicy *DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicy `field:"optional" json:"egressNetworkPolicy" yaml:"egressNetworkPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_room#region DataDatabricksCleanRoomsCleanRoom#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

