// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlquery


type SqlQueryParameterDatetimeRange struct {
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/sql_query#range SqlQuery#range}
	Range *SqlQueryParameterDatetimeRangeRange `field:"optional" json:"range" yaml:"range"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/sql_query#value SqlQuery#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

