package job


type JobJobClusterNewClusterInitScripts struct {
	// abfss block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/job#abfss Job#abfss}
	Abfss *JobJobClusterNewClusterInitScriptsAbfss `field:"optional" json:"abfss" yaml:"abfss"`
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/job#dbfs Job#dbfs}
	Dbfs *JobJobClusterNewClusterInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/job#file Job#file}
	File *JobJobClusterNewClusterInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/job#gcs Job#gcs}
	Gcs *JobJobClusterNewClusterInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/job#s3 Job#s3}
	S3 *JobJobClusterNewClusterInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
	// workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/job#workspace Job#workspace}
	Workspace *JobJobClusterNewClusterInitScriptsWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}

