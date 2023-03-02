package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsJobClusterNewClusterDockerImage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#url DataDatabricksJob#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#basic_auth DataDatabricksJob#basic_auth}
	BasicAuth *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterDockerImageBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

