package job


type JobTaskNewClusterDockerImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/job#url Job#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/job#basic_auth Job#basic_auth}
	BasicAuth *JobTaskNewClusterDockerImageBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

