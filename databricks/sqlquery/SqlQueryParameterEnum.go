// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlquery


type SqlQueryParameterEnum struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/sql_query#options SqlQuery#options}.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
	// multiple block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/sql_query#multiple SqlQuery#multiple}
	Multiple *SqlQueryParameterEnumMultiple `field:"optional" json:"multiple" yaml:"multiple"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/sql_query#value SqlQuery#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/sql_query#values SqlQuery#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

