package datadatabrickscluster


type DataDatabricksClusterClusterInfoDockerImageBasicAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/data-sources/cluster#password DataDatabricksCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/data-sources/cluster#username DataDatabricksCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

