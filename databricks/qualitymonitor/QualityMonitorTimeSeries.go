// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitor


type QualityMonitorTimeSeries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/quality_monitor#granularities QualityMonitor#granularities}.
	Granularities *[]*string `field:"required" json:"granularities" yaml:"granularities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/quality_monitor#timestamp_col QualityMonitor#timestamp_col}.
	TimestampCol *string `field:"required" json:"timestampCol" yaml:"timestampCol"`
}

