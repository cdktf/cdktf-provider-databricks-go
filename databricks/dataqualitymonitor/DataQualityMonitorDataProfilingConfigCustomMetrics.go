// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataqualitymonitor


type DataQualityMonitorDataProfilingConfigCustomMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/data_quality_monitor#definition DataQualityMonitor#definition}.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/data_quality_monitor#input_columns DataQualityMonitor#input_columns}.
	InputColumns *[]*string `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/data_quality_monitor#name DataQualityMonitor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/data_quality_monitor#output_data_type DataQualityMonitor#output_data_type}.
	OutputDataType *string `field:"required" json:"outputDataType" yaml:"outputDataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/data_quality_monitor#type DataQualityMonitor#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

