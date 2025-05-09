// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobJobCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/job#job_cluster_key Job#job_cluster_key}.
	JobClusterKey *string `field:"required" json:"jobClusterKey" yaml:"jobClusterKey"`
	// new_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/job#new_cluster Job#new_cluster}
	NewCluster *JobJobClusterNewCluster `field:"required" json:"newCluster" yaml:"newCluster"`
}

