// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoClusterLogConf struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/cluster#dbfs DataDatabricksCluster#dbfs}
	Dbfs *DataDatabricksClusterClusterInfoClusterLogConfDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/cluster#s3 DataDatabricksCluster#s3}
	S3 *DataDatabricksClusterClusterInfoClusterLogConfS3 `field:"optional" json:"s3" yaml:"s3"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/cluster#volumes DataDatabricksCluster#volumes}
	Volumes *DataDatabricksClusterClusterInfoClusterLogConfVolumes `field:"optional" json:"volumes" yaml:"volumes"`
}

