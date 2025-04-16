// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query


type QueryParameterEnumValueMultiValuesOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/query#prefix Query#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/query#separator Query#separator}.
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/query#suffix Query#suffix}.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

