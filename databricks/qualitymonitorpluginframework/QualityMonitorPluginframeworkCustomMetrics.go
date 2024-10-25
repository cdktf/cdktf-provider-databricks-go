// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorpluginframework


type QualityMonitorPluginframeworkCustomMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#definition QualityMonitorPluginframework#definition}.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#input_columns QualityMonitorPluginframework#input_columns}.
	InputColumns *[]*string `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#name QualityMonitorPluginframework#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#output_data_type QualityMonitorPluginframework#output_data_type}.
	OutputDataType *string `field:"required" json:"outputDataType" yaml:"outputDataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#type QualityMonitorPluginframework#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

