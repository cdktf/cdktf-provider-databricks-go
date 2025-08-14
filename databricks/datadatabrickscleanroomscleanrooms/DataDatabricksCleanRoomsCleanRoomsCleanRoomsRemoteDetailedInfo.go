// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomscleanrooms


type DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/data-sources/clean_rooms_clean_rooms#cloud_vendor DataDatabricksCleanRoomsCleanRooms#cloud_vendor}.
	CloudVendor *string `field:"optional" json:"cloudVendor" yaml:"cloudVendor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/data-sources/clean_rooms_clean_rooms#collaborators DataDatabricksCleanRoomsCleanRooms#collaborators}.
	Collaborators interface{} `field:"optional" json:"collaborators" yaml:"collaborators"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/data-sources/clean_rooms_clean_rooms#egress_network_policy DataDatabricksCleanRoomsCleanRooms#egress_network_policy}.
	EgressNetworkPolicy *DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicy `field:"optional" json:"egressNetworkPolicy" yaml:"egressNetworkPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/data-sources/clean_rooms_clean_rooms#region DataDatabricksCleanRoomsCleanRooms#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

