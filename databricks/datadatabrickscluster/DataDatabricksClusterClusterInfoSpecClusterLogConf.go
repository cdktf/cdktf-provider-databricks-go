// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoSpecClusterLogConf struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/data-sources/cluster#dbfs DataDatabricksCluster#dbfs}
	Dbfs *DataDatabricksClusterClusterInfoSpecClusterLogConfDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/data-sources/cluster#s3 DataDatabricksCluster#s3}
	S3 *DataDatabricksClusterClusterInfoSpecClusterLogConfS3 `field:"optional" json:"s3" yaml:"s3"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/data-sources/cluster#volumes DataDatabricksCluster#volumes}
	Volumes *DataDatabricksClusterClusterInfoSpecClusterLogConfVolumes `field:"optional" json:"volumes" yaml:"volumes"`
}

