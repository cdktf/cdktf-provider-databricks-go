package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTriggerFileArrival struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#url DataDatabricksJob#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#min_time_between_trigger_seconds DataDatabricksJob#min_time_between_trigger_seconds}.
	MinTimeBetweenTriggerSeconds *float64 `field:"optional" json:"minTimeBetweenTriggerSeconds" yaml:"minTimeBetweenTriggerSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#wait_after_last_change_seconds DataDatabricksJob#wait_after_last_change_seconds}.
	WaitAfterLastChangeSeconds *float64 `field:"optional" json:"waitAfterLastChangeSeconds" yaml:"waitAfterLastChangeSeconds"`
}

