// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoDiskSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/data-sources/instance_pool#disk_count DataDatabricksInstancePool#disk_count}.
	DiskCount *float64 `field:"optional" json:"diskCount" yaml:"diskCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/data-sources/instance_pool#disk_size DataDatabricksInstancePool#disk_size}.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// disk_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/data-sources/instance_pool#disk_type DataDatabricksInstancePool#disk_type}
	DiskType *DataDatabricksInstancePoolPoolInfoDiskSpecDiskType `field:"optional" json:"diskType" yaml:"diskType"`
}

