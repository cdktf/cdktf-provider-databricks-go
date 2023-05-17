package pipeline


type PipelineClusterClusterLogConf struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/pipeline#dbfs Pipeline#dbfs}
	Dbfs *PipelineClusterClusterLogConfDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/pipeline#s3 Pipeline#s3}
	S3 *PipelineClusterClusterLogConfS3 `field:"optional" json:"s3" yaml:"s3"`
}

