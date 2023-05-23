package job


type JobPipelineTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/job#pipeline_id Job#pipeline_id}.
	PipelineId *string `field:"required" json:"pipelineId" yaml:"pipelineId"`
}

