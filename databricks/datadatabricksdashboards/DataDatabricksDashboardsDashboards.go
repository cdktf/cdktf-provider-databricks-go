// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdashboards


type DataDatabricksDashboardsDashboards struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/dashboards#display_name DataDatabricksDashboards#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/dashboards#serialized_dashboard DataDatabricksDashboards#serialized_dashboard}.
	SerializedDashboard *string `field:"optional" json:"serializedDashboard" yaml:"serializedDashboard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/dashboards#warehouse_id DataDatabricksDashboards#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

