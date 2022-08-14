// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type InstancePoolInstancePoolFleetAttributesFleetSpotOption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/instance_pool#allocation_strategy InstancePool#allocation_strategy}.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/instance_pool#instance_pools_to_use_count InstancePool#instance_pools_to_use_count}.
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
}

