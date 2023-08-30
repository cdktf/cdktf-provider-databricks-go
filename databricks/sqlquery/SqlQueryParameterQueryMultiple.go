// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlquery


type SqlQueryParameterQueryMultiple struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/sql_query#prefix SqlQuery#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/sql_query#separator SqlQuery#separator}.
	Separator *string `field:"required" json:"separator" yaml:"separator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/sql_query#suffix SqlQuery#suffix}.
	Suffix *string `field:"required" json:"suffix" yaml:"suffix"`
}

