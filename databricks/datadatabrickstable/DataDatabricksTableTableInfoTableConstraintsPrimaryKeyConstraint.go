// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable


type DataDatabricksTableTableInfoTableConstraintsPrimaryKeyConstraint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.45.0/docs/data-sources/table#child_columns DataDatabricksTable#child_columns}.
	ChildColumns *[]*string `field:"required" json:"childColumns" yaml:"childColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.45.0/docs/data-sources/table#name DataDatabricksTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

