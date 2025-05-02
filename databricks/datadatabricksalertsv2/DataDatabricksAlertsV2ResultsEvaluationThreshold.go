// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertsv2


type DataDatabricksAlertsV2ResultsEvaluationThreshold struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.76.0/docs/data-sources/alerts_v2#column DataDatabricksAlertsV2#column}.
	Column *DataDatabricksAlertsV2ResultsEvaluationThresholdColumn `field:"optional" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.76.0/docs/data-sources/alerts_v2#value DataDatabricksAlertsV2#value}.
	Value *DataDatabricksAlertsV2ResultsEvaluationThresholdValue `field:"optional" json:"value" yaml:"value"`
}

