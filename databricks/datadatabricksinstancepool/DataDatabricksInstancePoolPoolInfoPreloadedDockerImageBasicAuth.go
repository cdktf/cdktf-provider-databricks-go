package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoPreloadedDockerImageBasicAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/data-sources/instance_pool#password DataDatabricksInstancePool#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/data-sources/instance_pool#username DataDatabricksInstancePool#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

