// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakehousemonitor


type LakehouseMonitorSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/lakehouse_monitor#quartz_cron_expression LakehouseMonitor#quartz_cron_expression}.
	QuartzCronExpression *string `field:"required" json:"quartzCronExpression" yaml:"quartzCronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/lakehouse_monitor#timezone_id LakehouseMonitor#timezone_id}.
	TimezoneId *string `field:"required" json:"timezoneId" yaml:"timezoneId"`
}

