// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssqlwarehouse


type DataDatabricksSqlWarehouseOdbcParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.27.0/docs/data-sources/sql_warehouse#path DataDatabricksSqlWarehouse#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.27.0/docs/data-sources/sql_warehouse#port DataDatabricksSqlWarehouse#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.27.0/docs/data-sources/sql_warehouse#protocol DataDatabricksSqlWarehouse#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.27.0/docs/data-sources/sql_warehouse#host DataDatabricksSqlWarehouse#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.27.0/docs/data-sources/sql_warehouse#hostname DataDatabricksSqlWarehouse#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}

