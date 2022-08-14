// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksClusterClusterInfoInitScripts struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#dbfs DataDatabricksCluster#dbfs}
	Dbfs *DataDatabricksClusterClusterInfoInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#s3 DataDatabricksCluster#s3}
	S3 *DataDatabricksClusterClusterInfoInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
}

