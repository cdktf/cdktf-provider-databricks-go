// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsJobCluster struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#job_cluster_key DataDatabricksJob#job_cluster_key}.
	JobClusterKey *string `field:"optional" json:"jobClusterKey" yaml:"jobClusterKey"`
	// new_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#new_cluster DataDatabricksJob#new_cluster}
	NewCluster *DataDatabricksJobJobSettingsSettingsJobClusterNewCluster `field:"optional" json:"newCluster" yaml:"newCluster"`
}

