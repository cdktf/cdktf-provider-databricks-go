// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTrigger struct {
	// file_arrival block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.0/docs/data-sources/job#file_arrival DataDatabricksJob#file_arrival}
	FileArrival *DataDatabricksJobJobSettingsSettingsTriggerFileArrival `field:"required" json:"fileArrival" yaml:"fileArrival"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.0/docs/data-sources/job#pause_status DataDatabricksJob#pause_status}.
	PauseStatus *string `field:"optional" json:"pauseStatus" yaml:"pauseStatus"`
}

