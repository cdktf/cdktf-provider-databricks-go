// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertv2


type AlertV2Evaluation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/alert_v2#comparison_operator AlertV2#comparison_operator}.
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/alert_v2#empty_result_state AlertV2#empty_result_state}.
	EmptyResultState *string `field:"optional" json:"emptyResultState" yaml:"emptyResultState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/alert_v2#notification AlertV2#notification}.
	Notification *AlertV2EvaluationNotification `field:"optional" json:"notification" yaml:"notification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/alert_v2#source AlertV2#source}.
	Source *AlertV2EvaluationSource `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/alert_v2#threshold AlertV2#threshold}.
	Threshold *AlertV2EvaluationThreshold `field:"optional" json:"threshold" yaml:"threshold"`
}

