// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsNewClusterInitScripts struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#dbfs DataDatabricksJob#dbfs}
	Dbfs *DataDatabricksJobJobSettingsSettingsNewClusterInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#file DataDatabricksJob#file}
	File *DataDatabricksJobJobSettingsSettingsNewClusterInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#gcs DataDatabricksJob#gcs}
	Gcs *DataDatabricksJobJobSettingsSettingsNewClusterInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#s3 DataDatabricksJob#s3}
	S3 *DataDatabricksJobJobSettingsSettingsNewClusterInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
}

