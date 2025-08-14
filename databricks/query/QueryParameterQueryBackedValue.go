// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query


type QueryParameterQueryBackedValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/query#query_id Query#query_id}.
	QueryId *string `field:"required" json:"queryId" yaml:"queryId"`
	// multi_values_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/query#multi_values_options Query#multi_values_options}
	MultiValuesOptions *QueryParameterQueryBackedValueMultiValuesOptions `field:"optional" json:"multiValuesOptions" yaml:"multiValuesOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/query#values Query#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

