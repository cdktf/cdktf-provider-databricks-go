// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobNewClusterDockerImageBasicAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#password Job#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#username Job#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

