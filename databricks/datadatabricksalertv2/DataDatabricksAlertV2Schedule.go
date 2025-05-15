// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertv2


type DataDatabricksAlertV2Schedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/alert_v2#pause_status DataDatabricksAlertV2#pause_status}.
	PauseStatus *string `field:"optional" json:"pauseStatus" yaml:"pauseStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/alert_v2#quartz_cron_schedule DataDatabricksAlertV2#quartz_cron_schedule}.
	QuartzCronSchedule *string `field:"optional" json:"quartzCronSchedule" yaml:"quartzCronSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/alert_v2#timezone_id DataDatabricksAlertV2#timezone_id}.
	TimezoneId *string `field:"optional" json:"timezoneId" yaml:"timezoneId"`
}

