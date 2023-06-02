package job


type JobWebhookNotifications struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/job#on_failure Job#on_failure}
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_start block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/job#on_start Job#on_start}
	OnStart interface{} `field:"optional" json:"onStart" yaml:"onStart"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/job#on_success Job#on_success}
	OnSuccess interface{} `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

