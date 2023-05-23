package pipeline


type PipelineNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/pipeline#alerts Pipeline#alerts}.
	Alerts *[]*string `field:"required" json:"alerts" yaml:"alerts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/pipeline#email_recipients Pipeline#email_recipients}.
	EmailRecipients *[]*string `field:"required" json:"emailRecipients" yaml:"emailRecipients"`
}

