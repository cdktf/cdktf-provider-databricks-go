// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertv2


type AlertV2Schedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/alert_v2#pause_status AlertV2#pause_status}.
	PauseStatus *string `field:"optional" json:"pauseStatus" yaml:"pauseStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/alert_v2#quartz_cron_schedule AlertV2#quartz_cron_schedule}.
	QuartzCronSchedule *string `field:"optional" json:"quartzCronSchedule" yaml:"quartzCronSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/alert_v2#timezone_id AlertV2#timezone_id}.
	TimezoneId *string `field:"optional" json:"timezoneId" yaml:"timezoneId"`
}

