// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertsv2


type DataDatabricksAlertsV2AlertsEvaluation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/alerts_v2#comparison_operator DataDatabricksAlertsV2#comparison_operator}.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/alerts_v2#source DataDatabricksAlertsV2#source}.
	Source *DataDatabricksAlertsV2AlertsEvaluationSource `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/alerts_v2#empty_result_state DataDatabricksAlertsV2#empty_result_state}.
	EmptyResultState *string `field:"optional" json:"emptyResultState" yaml:"emptyResultState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/alerts_v2#notification DataDatabricksAlertsV2#notification}.
	Notification *DataDatabricksAlertsV2AlertsEvaluationNotification `field:"optional" json:"notification" yaml:"notification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/alerts_v2#threshold DataDatabricksAlertsV2#threshold}.
	Threshold *DataDatabricksAlertsV2AlertsEvaluationThreshold `field:"optional" json:"threshold" yaml:"threshold"`
}

