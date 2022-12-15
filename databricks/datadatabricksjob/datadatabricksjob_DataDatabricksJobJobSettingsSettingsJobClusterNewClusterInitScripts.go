package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScripts struct {
	// abfss block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#abfss DataDatabricksJob#abfss}
	Abfss *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsAbfss `field:"optional" json:"abfss" yaml:"abfss"`
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#dbfs DataDatabricksJob#dbfs}
	Dbfs *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#file DataDatabricksJob#file}
	File *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#gcs DataDatabricksJob#gcs}
	Gcs *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#s3 DataDatabricksJob#s3}
	S3 *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
}

