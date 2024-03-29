// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakehousemonitor


type LakehouseMonitorSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#pause_status LakehouseMonitor#pause_status}.
	PauseStatus *string `field:"optional" json:"pauseStatus" yaml:"pauseStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#quartz_cron_expression LakehouseMonitor#quartz_cron_expression}.
	QuartzCronExpression *string `field:"optional" json:"quartzCronExpression" yaml:"quartzCronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#timezone_id LakehouseMonitor#timezone_id}.
	TimezoneId *string `field:"optional" json:"timezoneId" yaml:"timezoneId"`
}

