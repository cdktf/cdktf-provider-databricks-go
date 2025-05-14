// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertv2


type DataDatabricksAlertV2EvaluationThreshold struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/alert_v2#column DataDatabricksAlertV2#column}.
	Column *DataDatabricksAlertV2EvaluationThresholdColumn `field:"optional" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/alert_v2#value DataDatabricksAlertV2#value}.
	Value *DataDatabricksAlertV2EvaluationThresholdValue `field:"optional" json:"value" yaml:"value"`
}

