// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorpluginframework


type QualityMonitorPluginframeworkSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/quality_monitor_pluginframework#quartz_cron_expression QualityMonitorPluginframework#quartz_cron_expression}.
	QuartzCronExpression *string `field:"required" json:"quartzCronExpression" yaml:"quartzCronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/quality_monitor_pluginframework#timezone_id QualityMonitorPluginframework#timezone_id}.
	TimezoneId *string `field:"required" json:"timezoneId" yaml:"timezoneId"`
}

