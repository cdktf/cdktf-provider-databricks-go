package job


type JobTrigger struct {
	// file_arrival block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#file_arrival Job#file_arrival}
	FileArrival *JobTriggerFileArrival `field:"required" json:"fileArrival" yaml:"fileArrival"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#pause_status Job#pause_status}.
	PauseStatus *string `field:"optional" json:"pauseStatus" yaml:"pauseStatus"`
}

