package job


type JobTaskSqlTaskAlert struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.21.0/docs/resources/job#alert_id Job#alert_id}.
	AlertId *string `field:"required" json:"alertId" yaml:"alertId"`
	// subscriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.21.0/docs/resources/job#subscriptions Job#subscriptions}
	Subscriptions interface{} `field:"required" json:"subscriptions" yaml:"subscriptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.21.0/docs/resources/job#pause_subscriptions Job#pause_subscriptions}.
	PauseSubscriptions interface{} `field:"optional" json:"pauseSubscriptions" yaml:"pauseSubscriptions"`
}

