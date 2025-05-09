// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssqlwarehouse


type DataDatabricksSqlWarehouseOdbcParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/sql_warehouse#hostname DataDatabricksSqlWarehouse#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/sql_warehouse#path DataDatabricksSqlWarehouse#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/sql_warehouse#port DataDatabricksSqlWarehouse#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/sql_warehouse#protocol DataDatabricksSqlWarehouse#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

