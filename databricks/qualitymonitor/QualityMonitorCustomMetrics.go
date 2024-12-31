// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitor


type QualityMonitorCustomMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/quality_monitor#definition QualityMonitor#definition}.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/quality_monitor#input_columns QualityMonitor#input_columns}.
	InputColumns *[]*string `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/quality_monitor#name QualityMonitor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/quality_monitor#output_data_type QualityMonitor#output_data_type}.
	OutputDataType *string `field:"required" json:"outputDataType" yaml:"outputDataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/quality_monitor#type QualityMonitor#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

