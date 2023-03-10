package datadatabrickscluster


type DataDatabricksClusterClusterInfoInitScripts struct {
	// abfss block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#abfss DataDatabricksCluster#abfss}
	Abfss *DataDatabricksClusterClusterInfoInitScriptsAbfss `field:"optional" json:"abfss" yaml:"abfss"`
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#dbfs DataDatabricksCluster#dbfs}
	Dbfs *DataDatabricksClusterClusterInfoInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#file DataDatabricksCluster#file}
	File *DataDatabricksClusterClusterInfoInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#gcs DataDatabricksCluster#gcs}
	Gcs *DataDatabricksClusterClusterInfoInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#s3 DataDatabricksCluster#s3}
	S3 *DataDatabricksClusterClusterInfoInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
}

