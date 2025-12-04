// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskDashboardTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/job#dashboard_id DataDatabricksJob#dashboard_id}.
	DashboardId *string `field:"optional" json:"dashboardId" yaml:"dashboardId"`
	// subscription block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/job#subscription DataDatabricksJob#subscription}
	Subscription *DataDatabricksJobJobSettingsSettingsTaskDashboardTaskSubscription `field:"optional" json:"subscription" yaml:"subscription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/job#warehouse_id DataDatabricksJob#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

