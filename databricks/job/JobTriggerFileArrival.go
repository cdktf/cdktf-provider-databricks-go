package job


type JobTriggerFileArrival struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#url Job#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#min_time_between_trigger_seconds Job#min_time_between_trigger_seconds}.
	MinTimeBetweenTriggerSeconds *float64 `field:"optional" json:"minTimeBetweenTriggerSeconds" yaml:"minTimeBetweenTriggerSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#wait_after_last_change_seconds Job#wait_after_last_change_seconds}.
	WaitAfterLastChangeSeconds *float64 `field:"optional" json:"waitAfterLastChangeSeconds" yaml:"waitAfterLastChangeSeconds"`
}
