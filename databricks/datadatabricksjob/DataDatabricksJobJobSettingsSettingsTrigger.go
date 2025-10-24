// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTrigger struct {
	// file_arrival block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/job#file_arrival DataDatabricksJob#file_arrival}
	FileArrival *DataDatabricksJobJobSettingsSettingsTriggerFileArrival `field:"optional" json:"fileArrival" yaml:"fileArrival"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/job#pause_status DataDatabricksJob#pause_status}.
	PauseStatus *string `field:"optional" json:"pauseStatus" yaml:"pauseStatus"`
	// periodic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/job#periodic DataDatabricksJob#periodic}
	Periodic *DataDatabricksJobJobSettingsSettingsTriggerPeriodic `field:"optional" json:"periodic" yaml:"periodic"`
	// table_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/job#table_update DataDatabricksJob#table_update}
	TableUpdate *DataDatabricksJobJobSettingsSettingsTriggerTableUpdate `field:"optional" json:"tableUpdate" yaml:"tableUpdate"`
}

