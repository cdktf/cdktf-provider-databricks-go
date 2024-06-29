// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskNotebookTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.2/docs/resources/job#notebook_path Job#notebook_path}.
	NotebookPath *string `field:"required" json:"notebookPath" yaml:"notebookPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.2/docs/resources/job#base_parameters Job#base_parameters}.
	BaseParameters *map[string]*string `field:"optional" json:"baseParameters" yaml:"baseParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.2/docs/resources/job#source Job#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.2/docs/resources/job#warehouse_id Job#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

