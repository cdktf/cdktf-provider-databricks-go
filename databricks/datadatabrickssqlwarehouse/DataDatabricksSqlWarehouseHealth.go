// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssqlwarehouse


type DataDatabricksSqlWarehouseHealth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/sql_warehouse#details DataDatabricksSqlWarehouse#details}.
	Details *string `field:"optional" json:"details" yaml:"details"`
	// failure_reason block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/sql_warehouse#failure_reason DataDatabricksSqlWarehouse#failure_reason}
	FailureReason *DataDatabricksSqlWarehouseHealthFailureReason `field:"optional" json:"failureReason" yaml:"failureReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/sql_warehouse#message DataDatabricksSqlWarehouse#message}.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/sql_warehouse#status DataDatabricksSqlWarehouse#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/sql_warehouse#summary DataDatabricksSqlWarehouse#summary}.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
}

