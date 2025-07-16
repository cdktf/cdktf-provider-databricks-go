// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertv2


type AlertV2EvaluationThresholdValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/alert_v2#bool_value AlertV2#bool_value}.
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/alert_v2#double_value AlertV2#double_value}.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/alert_v2#string_value AlertV2#string_value}.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

