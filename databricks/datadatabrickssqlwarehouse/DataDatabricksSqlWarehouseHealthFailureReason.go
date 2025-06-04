// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssqlwarehouse


type DataDatabricksSqlWarehouseHealthFailureReason struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/data-sources/sql_warehouse#code DataDatabricksSqlWarehouse#code}.
	Code *string `field:"optional" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/data-sources/sql_warehouse#parameters DataDatabricksSqlWarehouse#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/data-sources/sql_warehouse#type DataDatabricksSqlWarehouse#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

