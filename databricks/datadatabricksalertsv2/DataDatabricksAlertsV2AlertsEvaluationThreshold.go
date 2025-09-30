// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertsv2


type DataDatabricksAlertsV2AlertsEvaluationThreshold struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/alerts_v2#column DataDatabricksAlertsV2#column}.
	Column *DataDatabricksAlertsV2AlertsEvaluationThresholdColumn `field:"optional" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/alerts_v2#value DataDatabricksAlertsV2#value}.
	Value *DataDatabricksAlertsV2AlertsEvaluationThresholdValue `field:"optional" json:"value" yaml:"value"`
}

