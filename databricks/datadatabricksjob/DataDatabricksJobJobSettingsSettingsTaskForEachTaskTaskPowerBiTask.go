// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPowerBiTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#connection_resource_name DataDatabricksJob#connection_resource_name}.
	ConnectionResourceName *string `field:"optional" json:"connectionResourceName" yaml:"connectionResourceName"`
	// power_bi_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#power_bi_model DataDatabricksJob#power_bi_model}
	PowerBiModel *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPowerBiTaskPowerBiModel `field:"optional" json:"powerBiModel" yaml:"powerBiModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#refresh_after_update DataDatabricksJob#refresh_after_update}.
	RefreshAfterUpdate interface{} `field:"optional" json:"refreshAfterUpdate" yaml:"refreshAfterUpdate"`
	// tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#tables DataDatabricksJob#tables}
	Tables interface{} `field:"optional" json:"tables" yaml:"tables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#warehouse_id DataDatabricksJob#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

