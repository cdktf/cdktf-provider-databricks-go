package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsJobClusterNewClusterDockerImageBasicAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#password DataDatabricksJob#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#username DataDatabricksJob#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}
