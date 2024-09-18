// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable


type DataDatabricksTableTableInfoColumnsMask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/table#function_name DataDatabricksTable#function_name}.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/table#using_column_names DataDatabricksTable#using_column_names}.
	UsingColumnNames *[]*string `field:"optional" json:"usingColumnNames" yaml:"usingColumnNames"`
}

