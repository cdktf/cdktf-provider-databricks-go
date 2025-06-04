// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query


type QueryParameterEnumValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/query#enum_options Query#enum_options}.
	EnumOptions *string `field:"optional" json:"enumOptions" yaml:"enumOptions"`
	// multi_values_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/query#multi_values_options Query#multi_values_options}
	MultiValuesOptions *QueryParameterEnumValueMultiValuesOptions `field:"optional" json:"multiValuesOptions" yaml:"multiValuesOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/query#values Query#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

