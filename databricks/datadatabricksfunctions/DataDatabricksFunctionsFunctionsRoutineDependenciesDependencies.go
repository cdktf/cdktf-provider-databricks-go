// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfunctions


type DataDatabricksFunctionsFunctionsRoutineDependenciesDependencies struct {
	// function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/functions#function DataDatabricksFunctions#function}
	Function interface{} `field:"optional" json:"function" yaml:"function"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/functions#table DataDatabricksFunctions#table}
	Table interface{} `field:"optional" json:"table" yaml:"table"`
}
