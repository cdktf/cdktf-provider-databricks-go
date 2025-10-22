// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertv2


type DataDatabricksAlertV2EvaluationThresholdValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/alert_v2#bool_value DataDatabricksAlertV2#bool_value}.
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/alert_v2#double_value DataDatabricksAlertV2#double_value}.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/alert_v2#string_value DataDatabricksAlertV2#string_value}.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

