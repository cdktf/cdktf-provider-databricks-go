package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTriggerFileArrival struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/data-sources/job#url DataDatabricksJob#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/data-sources/job#min_time_between_triggers_seconds DataDatabricksJob#min_time_between_triggers_seconds}.
	MinTimeBetweenTriggersSeconds *float64 `field:"optional" json:"minTimeBetweenTriggersSeconds" yaml:"minTimeBetweenTriggersSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/data-sources/job#wait_after_last_change_seconds DataDatabricksJob#wait_after_last_change_seconds}.
	WaitAfterLastChangeSeconds *float64 `field:"optional" json:"waitAfterLastChangeSeconds" yaml:"waitAfterLastChangeSeconds"`
}

