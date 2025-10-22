// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskDashboardTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/job#dashboard_id Job#dashboard_id}.
	DashboardId *string `field:"optional" json:"dashboardId" yaml:"dashboardId"`
	// subscription block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/job#subscription Job#subscription}
	Subscription *JobTaskForEachTaskTaskDashboardTaskSubscription `field:"optional" json:"subscription" yaml:"subscription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/job#warehouse_id Job#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

