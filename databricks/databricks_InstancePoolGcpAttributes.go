// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type InstancePoolGcpAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/instance_pool#gcp_availability InstancePool#gcp_availability}.
	GcpAvailability *string `field:"optional" json:"gcpAvailability" yaml:"gcpAvailability"`
}

