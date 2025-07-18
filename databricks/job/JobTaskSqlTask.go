// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskSqlTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#warehouse_id Job#warehouse_id}.
	WarehouseId *string `field:"required" json:"warehouseId" yaml:"warehouseId"`
	// alert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#alert Job#alert}
	Alert *JobTaskSqlTaskAlert `field:"optional" json:"alert" yaml:"alert"`
	// dashboard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#dashboard Job#dashboard}
	Dashboard *JobTaskSqlTaskDashboard `field:"optional" json:"dashboard" yaml:"dashboard"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#file Job#file}
	File *JobTaskSqlTaskFile `field:"optional" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#parameters Job#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/job#query Job#query}
	Query *JobTaskSqlTaskQuery `field:"optional" json:"query" yaml:"query"`
}

