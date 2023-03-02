package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverride struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#availability_zone DataDatabricksInstancePool#availability_zone}.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#instance_type DataDatabricksInstancePool#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

