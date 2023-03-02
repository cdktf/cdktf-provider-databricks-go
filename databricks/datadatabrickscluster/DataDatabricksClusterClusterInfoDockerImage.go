package datadatabrickscluster


type DataDatabricksClusterClusterInfoDockerImage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#url DataDatabricksCluster#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#basic_auth DataDatabricksCluster#basic_auth}
	BasicAuth *DataDatabricksClusterClusterInfoDockerImageBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

