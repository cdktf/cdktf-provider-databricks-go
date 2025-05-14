// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alert


type AlertConditionThresholdValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/resources/alert#bool_value Alert#bool_value}.
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/resources/alert#double_value Alert#double_value}.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/resources/alert#string_value Alert#string_value}.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

