// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorpluginframework


type QualityMonitorPluginframeworkTimeSeries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.56.0/docs/resources/quality_monitor_pluginframework#granularities QualityMonitorPluginframework#granularities}.
	Granularities *[]*string `field:"required" json:"granularities" yaml:"granularities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.56.0/docs/resources/quality_monitor_pluginframework#timestamp_col QualityMonitorPluginframework#timestamp_col}.
	TimestampCol *string `field:"required" json:"timestampCol" yaml:"timestampCol"`
}

