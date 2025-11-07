// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsJobCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/job#job_cluster_key DataDatabricksJob#job_cluster_key}.
	JobClusterKey *string `field:"required" json:"jobClusterKey" yaml:"jobClusterKey"`
	// new_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/data-sources/job#new_cluster DataDatabricksJob#new_cluster}
	NewCluster *DataDatabricksJobJobSettingsSettingsJobClusterNewCluster `field:"required" json:"newCluster" yaml:"newCluster"`
}

