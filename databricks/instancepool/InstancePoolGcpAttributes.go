package instancepool


type InstancePoolGcpAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/resources/instance_pool#gcp_availability InstancePool#gcp_availability}.
	GcpAvailability *string `field:"optional" json:"gcpAvailability" yaml:"gcpAvailability"`
}

