package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoPreloadedDockerImageBasicAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#password DataDatabricksInstancePool#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#username DataDatabricksInstancePool#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

