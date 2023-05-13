package job


type JobTaskDependsOn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.0/docs/resources/job#task_key Job#task_key}.
	TaskKey *string `field:"optional" json:"taskKey" yaml:"taskKey"`
}

