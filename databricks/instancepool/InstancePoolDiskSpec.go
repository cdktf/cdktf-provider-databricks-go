package instancepool


type InstancePoolDiskSpec struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/instance_pool#disk_count InstancePool#disk_count}.
	DiskCount *float64 `field:"optional" json:"diskCount" yaml:"diskCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/instance_pool#disk_size InstancePool#disk_size}.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// disk_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/instance_pool#disk_type InstancePool#disk_type}
	DiskType *InstancePoolDiskSpecDiskType `field:"optional" json:"diskType" yaml:"diskType"`
}
