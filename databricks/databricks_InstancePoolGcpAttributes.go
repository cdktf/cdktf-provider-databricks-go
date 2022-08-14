// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type InstancePoolGcpAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/instance_pool#availability InstancePool#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
}

