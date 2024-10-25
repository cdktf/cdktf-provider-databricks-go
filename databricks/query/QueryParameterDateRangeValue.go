// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query


type QueryParameterDateRangeValue struct {
	// date_range_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/query#date_range_value Query#date_range_value}
	DateRangeValue *QueryParameterDateRangeValueDateRangeValue `field:"optional" json:"dateRangeValue" yaml:"dateRangeValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/query#dynamic_date_range_value Query#dynamic_date_range_value}.
	DynamicDateRangeValue *string `field:"optional" json:"dynamicDateRangeValue" yaml:"dynamicDateRangeValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/query#precision Query#precision}.
	Precision *string `field:"optional" json:"precision" yaml:"precision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/query#start_day_of_week Query#start_day_of_week}.
	StartDayOfWeek *float64 `field:"optional" json:"startDayOfWeek" yaml:"startDayOfWeek"`
}

