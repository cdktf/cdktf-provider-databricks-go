// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlquery


type SqlQueryParameterQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/sql_query#query_id SqlQuery#query_id}.
	QueryId *string `field:"required" json:"queryId" yaml:"queryId"`
	// multiple block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/sql_query#multiple SqlQuery#multiple}
	Multiple *SqlQueryParameterQueryMultiple `field:"optional" json:"multiple" yaml:"multiple"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/sql_query#value SqlQuery#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/sql_query#values SqlQuery#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

