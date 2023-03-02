package job


type JobTaskSqlTaskAlert struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#alert_id Job#alert_id}.
	AlertId *string `field:"required" json:"alertId" yaml:"alertId"`
}

