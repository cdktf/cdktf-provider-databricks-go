// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobTaskPipelineTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#pipeline_id Job#pipeline_id}.
	PipelineId *string `field:"required" json:"pipelineId" yaml:"pipelineId"`
}

