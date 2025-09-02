// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query


type QueryParameterDateRangeValueDateRangeValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/query#end Query#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/query#start Query#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

