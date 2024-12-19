// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query


type QueryParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.0/docs/resources/query#name Query#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// date_range_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.0/docs/resources/query#date_range_value Query#date_range_value}
	DateRangeValue *QueryParameterDateRangeValue `field:"optional" json:"dateRangeValue" yaml:"dateRangeValue"`
	// date_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.0/docs/resources/query#date_value Query#date_value}
	DateValue *QueryParameterDateValue `field:"optional" json:"dateValue" yaml:"dateValue"`
	// enum_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.0/docs/resources/query#enum_value Query#enum_value}
	EnumValue *QueryParameterEnumValue `field:"optional" json:"enumValue" yaml:"enumValue"`
	// numeric_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.0/docs/resources/query#numeric_value Query#numeric_value}
	NumericValue *QueryParameterNumericValue `field:"optional" json:"numericValue" yaml:"numericValue"`
	// query_backed_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.0/docs/resources/query#query_backed_value Query#query_backed_value}
	QueryBackedValue *QueryParameterQueryBackedValue `field:"optional" json:"queryBackedValue" yaml:"queryBackedValue"`
	// text_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.0/docs/resources/query#text_value Query#text_value}
	TextValue *QueryParameterTextValue `field:"optional" json:"textValue" yaml:"textValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.0/docs/resources/query#title Query#title}.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

