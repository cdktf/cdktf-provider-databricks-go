package cluster


type ClusterInitScripts struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#dbfs Cluster#dbfs}
	Dbfs *ClusterInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#file Cluster#file}
	File *ClusterInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#gcs Cluster#gcs}
	Gcs *ClusterInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#s3 Cluster#s3}
	S3 *ClusterInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
}

