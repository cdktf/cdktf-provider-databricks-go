package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoPreloadedDockerImage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#url DataDatabricksInstancePool#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#basic_auth DataDatabricksInstancePool#basic_auth}
	BasicAuth *DataDatabricksInstancePoolPoolInfoPreloadedDockerImageBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

