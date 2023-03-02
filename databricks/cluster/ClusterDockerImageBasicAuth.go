package cluster


type ClusterDockerImageBasicAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#password Cluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#username Cluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

