package instancepool


type InstancePoolPreloadedDockerImageBasicAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/instance_pool#password InstancePool#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/instance_pool#username InstancePool#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

