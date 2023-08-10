package job


type JobTaskRunJobTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.23.0/docs/resources/job#job_id Job#job_id}.
	JobId *string `field:"required" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.23.0/docs/resources/job#job_parameters Job#job_parameters}.
	JobParameters *map[string]*string `field:"optional" json:"jobParameters" yaml:"jobParameters"`
}

