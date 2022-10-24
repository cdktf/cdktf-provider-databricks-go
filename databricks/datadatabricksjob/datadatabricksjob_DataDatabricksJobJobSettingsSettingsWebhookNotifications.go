package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsWebhookNotifications struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#on_failure DataDatabricksJob#on_failure}
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#on_start DataDatabricksJob#on_start}
	OnStart interface{} `field:"optional" json:"onStart" yaml:"onStart"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#on_success DataDatabricksJob#on_success}
	OnSuccess interface{} `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

