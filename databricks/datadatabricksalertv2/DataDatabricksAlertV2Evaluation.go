// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertv2


type DataDatabricksAlertV2Evaluation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/alert_v2#comparison_operator DataDatabricksAlertV2#comparison_operator}.
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/alert_v2#empty_result_state DataDatabricksAlertV2#empty_result_state}.
	EmptyResultState *string `field:"optional" json:"emptyResultState" yaml:"emptyResultState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/alert_v2#notification DataDatabricksAlertV2#notification}.
	Notification *DataDatabricksAlertV2EvaluationNotification `field:"optional" json:"notification" yaml:"notification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/alert_v2#source DataDatabricksAlertV2#source}.
	Source *DataDatabricksAlertV2EvaluationSource `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/alert_v2#threshold DataDatabricksAlertV2#threshold}.
	Threshold *DataDatabricksAlertV2EvaluationThreshold `field:"optional" json:"threshold" yaml:"threshold"`
}

