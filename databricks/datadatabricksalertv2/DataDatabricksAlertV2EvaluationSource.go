// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertv2


type DataDatabricksAlertV2EvaluationSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/alert_v2#name DataDatabricksAlertV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/alert_v2#aggregation DataDatabricksAlertV2#aggregation}.
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/alert_v2#display DataDatabricksAlertV2#display}.
	Display *string `field:"optional" json:"display" yaml:"display"`
}

