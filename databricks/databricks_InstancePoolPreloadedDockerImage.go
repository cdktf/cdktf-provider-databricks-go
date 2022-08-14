// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type InstancePoolPreloadedDockerImage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/instance_pool#url InstancePool#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/instance_pool#basic_auth InstancePool#basic_auth}
	BasicAuth *InstancePoolPreloadedDockerImageBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

