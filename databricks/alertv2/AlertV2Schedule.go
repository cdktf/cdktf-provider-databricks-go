// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertv2


type AlertV2Schedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/alert_v2#quartz_cron_schedule AlertV2#quartz_cron_schedule}.
	QuartzCronSchedule *string `field:"required" json:"quartzCronSchedule" yaml:"quartzCronSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/alert_v2#timezone_id AlertV2#timezone_id}.
	TimezoneId *string `field:"required" json:"timezoneId" yaml:"timezoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/alert_v2#pause_status AlertV2#pause_status}.
	PauseStatus *string `field:"optional" json:"pauseStatus" yaml:"pauseStatus"`
}

