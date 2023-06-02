package job


type JobTaskSqlTaskAlert struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/job#alert_id Job#alert_id}.
	AlertId *string `field:"required" json:"alertId" yaml:"alertId"`
}

