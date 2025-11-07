// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfunctions


type DataDatabricksFunctionsFunctionsRoutineDependenciesDependencies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/functions#connection DataDatabricksFunctions#connection}.
	Connection *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesConnection `field:"optional" json:"connection" yaml:"connection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/functions#credential DataDatabricksFunctions#credential}.
	Credential *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesCredential `field:"optional" json:"credential" yaml:"credential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/functions#function DataDatabricksFunctions#function}.
	Function *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesFunction `field:"optional" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/functions#table DataDatabricksFunctions#table}.
	Table *DataDatabricksFunctionsFunctionsRoutineDependenciesDependenciesTable `field:"optional" json:"table" yaml:"table"`
}

