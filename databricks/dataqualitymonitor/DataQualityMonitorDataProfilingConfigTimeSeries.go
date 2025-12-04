// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataqualitymonitor


type DataQualityMonitorDataProfilingConfigTimeSeries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/data_quality_monitor#granularities DataQualityMonitor#granularities}.
	Granularities *[]*string `field:"required" json:"granularities" yaml:"granularities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/data_quality_monitor#timestamp_column DataQualityMonitor#timestamp_column}.
	TimestampColumn *string `field:"required" json:"timestampColumn" yaml:"timestampColumn"`
}

