// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query


type QueryParameterDateValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/query#date_value Query#date_value}.
	DateValue *string `field:"optional" json:"dateValue" yaml:"dateValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/query#dynamic_date_value Query#dynamic_date_value}.
	DynamicDateValue *string `field:"optional" json:"dynamicDateValue" yaml:"dynamicDateValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/query#precision Query#precision}.
	Precision *string `field:"optional" json:"precision" yaml:"precision"`
}

