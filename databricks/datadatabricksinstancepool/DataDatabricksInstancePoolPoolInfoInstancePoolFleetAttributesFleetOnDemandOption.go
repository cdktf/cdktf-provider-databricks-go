package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/data-sources/instance_pool#allocation_strategy DataDatabricksInstancePool#allocation_strategy}.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/data-sources/instance_pool#instance_pools_to_use_count DataDatabricksInstancePool#instance_pools_to_use_count}.
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
}

