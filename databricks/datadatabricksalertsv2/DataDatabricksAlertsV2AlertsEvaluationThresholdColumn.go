// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertsv2


type DataDatabricksAlertsV2AlertsEvaluationThresholdColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/alerts_v2#name DataDatabricksAlertsV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/alerts_v2#aggregation DataDatabricksAlertsV2#aggregation}.
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/alerts_v2#display DataDatabricksAlertsV2#display}.
	Display *string `field:"optional" json:"display" yaml:"display"`
}

