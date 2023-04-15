package pipeline


type PipelineNotification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#alerts Pipeline#alerts}.
	Alerts *[]*string `field:"required" json:"alerts" yaml:"alerts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#email_recipients Pipeline#email_recipients}.
	EmailRecipients *[]*string `field:"required" json:"emailRecipients" yaml:"emailRecipients"`
}

