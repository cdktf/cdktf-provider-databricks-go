// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitors


type DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfigTimeSeries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#granularities DataDatabricksDataQualityMonitors#granularities}.
	Granularities *[]*string `field:"required" json:"granularities" yaml:"granularities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#timestamp_column DataDatabricksDataQualityMonitors#timestamp_column}.
	TimestampColumn *string `field:"required" json:"timestampColumn" yaml:"timestampColumn"`
}

