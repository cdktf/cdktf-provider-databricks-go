// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitor


type DataDatabricksDataQualityMonitorDataProfilingConfigCustomMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitor#definition DataDatabricksDataQualityMonitor#definition}.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitor#input_columns DataDatabricksDataQualityMonitor#input_columns}.
	InputColumns *[]*string `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitor#name DataDatabricksDataQualityMonitor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitor#output_data_type DataDatabricksDataQualityMonitor#output_data_type}.
	OutputDataType *string `field:"required" json:"outputDataType" yaml:"outputDataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitor#type DataDatabricksDataQualityMonitor#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

