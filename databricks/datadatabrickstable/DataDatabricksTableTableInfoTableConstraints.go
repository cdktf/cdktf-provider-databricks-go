// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable


type DataDatabricksTableTableInfoTableConstraints struct {
	// foreign_key_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.71.0/docs/data-sources/table#foreign_key_constraint DataDatabricksTable#foreign_key_constraint}
	ForeignKeyConstraint *DataDatabricksTableTableInfoTableConstraintsForeignKeyConstraint `field:"optional" json:"foreignKeyConstraint" yaml:"foreignKeyConstraint"`
	// named_table_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.71.0/docs/data-sources/table#named_table_constraint DataDatabricksTable#named_table_constraint}
	NamedTableConstraint *DataDatabricksTableTableInfoTableConstraintsNamedTableConstraint `field:"optional" json:"namedTableConstraint" yaml:"namedTableConstraint"`
	// primary_key_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.71.0/docs/data-sources/table#primary_key_constraint DataDatabricksTable#primary_key_constraint}
	PrimaryKeyConstraint *DataDatabricksTableTableInfoTableConstraintsPrimaryKeyConstraint `field:"optional" json:"primaryKeyConstraint" yaml:"primaryKeyConstraint"`
}

