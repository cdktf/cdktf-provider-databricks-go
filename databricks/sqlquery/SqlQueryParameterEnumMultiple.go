// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlquery


type SqlQueryParameterEnumMultiple struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/resources/sql_query#separator SqlQuery#separator}.
	Separator *string `field:"required" json:"separator" yaml:"separator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/resources/sql_query#prefix SqlQuery#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/resources/sql_query#suffix SqlQuery#suffix}.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

