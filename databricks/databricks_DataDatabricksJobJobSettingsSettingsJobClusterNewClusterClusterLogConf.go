// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsJobClusterNewClusterClusterLogConf struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#dbfs DataDatabricksJob#dbfs}
	Dbfs *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterClusterLogConfDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#s3 DataDatabricksJob#s3}
	S3 *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterClusterLogConfS3 `field:"optional" json:"s3" yaml:"s3"`
}

