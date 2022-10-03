package cluster


type ClusterClusterLogConf struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#dbfs Cluster#dbfs}
	Dbfs *ClusterClusterLogConfDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#s3 Cluster#s3}
	S3 *ClusterClusterLogConfS3 `field:"optional" json:"s3" yaml:"s3"`
}

