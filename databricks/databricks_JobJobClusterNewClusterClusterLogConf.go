// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobJobClusterNewClusterClusterLogConf struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#dbfs Job#dbfs}
	Dbfs *JobJobClusterNewClusterClusterLogConfDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#s3 Job#s3}
	S3 *JobJobClusterNewClusterClusterLogConfS3 `field:"optional" json:"s3" yaml:"s3"`
}

