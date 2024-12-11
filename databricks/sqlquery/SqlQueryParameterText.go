// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlquery


type SqlQueryParameterText struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.60.0/docs/resources/sql_query#value SqlQuery#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

