// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineTriggerCron struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/pipeline#quartz_cron_schedule Pipeline#quartz_cron_schedule}.
	QuartzCronSchedule *string `field:"optional" json:"quartzCronSchedule" yaml:"quartzCronSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/pipeline#timezone_id Pipeline#timezone_id}.
	TimezoneId *string `field:"optional" json:"timezoneId" yaml:"timezoneId"`
}

