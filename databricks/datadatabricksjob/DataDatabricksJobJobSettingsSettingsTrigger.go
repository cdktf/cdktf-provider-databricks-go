package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTrigger struct {
	// file_arrival block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#file_arrival DataDatabricksJob#file_arrival}
	FileArrival *DataDatabricksJobJobSettingsSettingsTriggerFileArrival `field:"required" json:"fileArrival" yaml:"fileArrival"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#pause_status DataDatabricksJob#pause_status}.
	PauseStatus *string `field:"optional" json:"pauseStatus" yaml:"pauseStatus"`
}

