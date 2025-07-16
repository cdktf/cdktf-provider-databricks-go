// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksonlinestores


type DataDatabricksOnlineStoresOnlineStores struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/online_stores#capacity DataDatabricksOnlineStores#capacity}.
	Capacity *string `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/online_stores#name DataDatabricksOnlineStores#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/online_stores#read_replica_count DataDatabricksOnlineStores#read_replica_count}.
	ReadReplicaCount *float64 `field:"optional" json:"readReplicaCount" yaml:"readReplicaCount"`
}

