// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitors


type DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfigCustomMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#definition DataDatabricksDataQualityMonitors#definition}.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#input_columns DataDatabricksDataQualityMonitors#input_columns}.
	InputColumns *[]*string `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#name DataDatabricksDataQualityMonitors#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#output_data_type DataDatabricksDataQualityMonitors#output_data_type}.
	OutputDataType *string `field:"required" json:"outputDataType" yaml:"outputDataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#type DataDatabricksDataQualityMonitors#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

