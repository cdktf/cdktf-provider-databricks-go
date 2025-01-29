// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable


type DataDatabricksTableTableInfoViewDependenciesDependencies struct {
	// function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/table#function DataDatabricksTable#function}
	Function *DataDatabricksTableTableInfoViewDependenciesDependenciesFunction `field:"optional" json:"function" yaml:"function"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/table#table DataDatabricksTable#table}
	Table *DataDatabricksTableTableInfoViewDependenciesDependenciesTable `field:"optional" json:"table" yaml:"table"`
}

