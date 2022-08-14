// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksClusterClusterInfoDockerImageBasicAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#password DataDatabricksCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#username DataDatabricksCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

