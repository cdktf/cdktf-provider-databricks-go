// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitor


type DataDatabricksDataQualityMonitorDataProfilingConfigSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/data_quality_monitor#quartz_cron_expression DataDatabricksDataQualityMonitor#quartz_cron_expression}.
	QuartzCronExpression *string `field:"required" json:"quartzCronExpression" yaml:"quartzCronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/data_quality_monitor#timezone_id DataDatabricksDataQualityMonitor#timezone_id}.
	TimezoneId *string `field:"required" json:"timezoneId" yaml:"timezoneId"`
}

