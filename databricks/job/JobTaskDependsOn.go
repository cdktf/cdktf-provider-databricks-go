package job


type JobTaskDependsOn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#task_key Job#task_key}.
	TaskKey *string `field:"optional" json:"taskKey" yaml:"taskKey"`
}

