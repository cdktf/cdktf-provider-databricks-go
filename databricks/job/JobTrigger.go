// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTrigger struct {
	// file_arrival block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.46.0/docs/resources/job#file_arrival Job#file_arrival}
	FileArrival *JobTriggerFileArrival `field:"optional" json:"fileArrival" yaml:"fileArrival"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.46.0/docs/resources/job#pause_status Job#pause_status}.
	PauseStatus *string `field:"optional" json:"pauseStatus" yaml:"pauseStatus"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.46.0/docs/resources/job#table Job#table}
	Table *JobTriggerTable `field:"optional" json:"table" yaml:"table"`
	// table_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.46.0/docs/resources/job#table_update Job#table_update}
	TableUpdate *JobTriggerTableUpdate `field:"optional" json:"tableUpdate" yaml:"tableUpdate"`
}

