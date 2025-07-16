// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskSqlTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/job#warehouse_id DataDatabricksJob#warehouse_id}.
	WarehouseId *string `field:"required" json:"warehouseId" yaml:"warehouseId"`
	// alert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/job#alert DataDatabricksJob#alert}
	Alert *DataDatabricksJobJobSettingsSettingsTaskSqlTaskAlert `field:"optional" json:"alert" yaml:"alert"`
	// dashboard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/job#dashboard DataDatabricksJob#dashboard}
	Dashboard *DataDatabricksJobJobSettingsSettingsTaskSqlTaskDashboard `field:"optional" json:"dashboard" yaml:"dashboard"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/job#file DataDatabricksJob#file}
	File *DataDatabricksJobJobSettingsSettingsTaskSqlTaskFile `field:"optional" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/job#parameters DataDatabricksJob#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/job#query DataDatabricksJob#query}
	Query *DataDatabricksJobJobSettingsSettingsTaskSqlTaskQuery `field:"optional" json:"query" yaml:"query"`
}

