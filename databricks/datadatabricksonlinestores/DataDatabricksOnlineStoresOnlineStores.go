// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksonlinestores


type DataDatabricksOnlineStoresOnlineStores struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/data-sources/online_stores#name DataDatabricksOnlineStores#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/data-sources/online_stores#capacity DataDatabricksOnlineStores#capacity}.
	Capacity *string `field:"optional" json:"capacity" yaml:"capacity"`
}

