package job


type JobTaskPipelineTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/job#pipeline_id Job#pipeline_id}.
	PipelineId *string `field:"required" json:"pipelineId" yaml:"pipelineId"`
}

