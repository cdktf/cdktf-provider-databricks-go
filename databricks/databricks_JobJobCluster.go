// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobJobCluster struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#job_cluster_key Job#job_cluster_key}.
	JobClusterKey *string `field:"optional" json:"jobClusterKey" yaml:"jobClusterKey"`
	// new_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#new_cluster Job#new_cluster}
	NewCluster *JobJobClusterNewCluster `field:"optional" json:"newCluster" yaml:"newCluster"`
}

