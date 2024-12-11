// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskHealthRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.60.0/docs/data-sources/job#metric DataDatabricksJob#metric}.
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.60.0/docs/data-sources/job#op DataDatabricksJob#op}.
	Op *string `field:"required" json:"op" yaml:"op"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.60.0/docs/data-sources/job#value DataDatabricksJob#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

