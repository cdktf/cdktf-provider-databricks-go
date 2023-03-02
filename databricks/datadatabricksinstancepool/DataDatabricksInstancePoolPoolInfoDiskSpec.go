package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoDiskSpec struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#disk_count DataDatabricksInstancePool#disk_count}.
	DiskCount *float64 `field:"optional" json:"diskCount" yaml:"diskCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#disk_size DataDatabricksInstancePool#disk_size}.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// disk_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#disk_type DataDatabricksInstancePool#disk_type}
	DiskType *DataDatabricksInstancePoolPoolInfoDiskSpecDiskType `field:"optional" json:"diskType" yaml:"diskType"`
}

