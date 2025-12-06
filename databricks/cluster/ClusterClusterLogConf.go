// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterClusterLogConf struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/cluster#dbfs Cluster#dbfs}
	Dbfs *ClusterClusterLogConfDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/cluster#s3 Cluster#s3}
	S3 *ClusterClusterLogConfS3 `field:"optional" json:"s3" yaml:"s3"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/cluster#volumes Cluster#volumes}
	Volumes *ClusterClusterLogConfVolumes `field:"optional" json:"volumes" yaml:"volumes"`
}

