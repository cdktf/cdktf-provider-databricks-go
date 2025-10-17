// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskPowerBiTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/job#connection_resource_name Job#connection_resource_name}.
	ConnectionResourceName *string `field:"optional" json:"connectionResourceName" yaml:"connectionResourceName"`
	// power_bi_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/job#power_bi_model Job#power_bi_model}
	PowerBiModel *JobTaskForEachTaskTaskPowerBiTaskPowerBiModel `field:"optional" json:"powerBiModel" yaml:"powerBiModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/job#refresh_after_update Job#refresh_after_update}.
	RefreshAfterUpdate interface{} `field:"optional" json:"refreshAfterUpdate" yaml:"refreshAfterUpdate"`
	// tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/job#tables Job#tables}
	Tables interface{} `field:"optional" json:"tables" yaml:"tables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/job#warehouse_id Job#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

