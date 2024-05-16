// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobJobClusterNewClusterCloneFrom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/resources/job#source_cluster_id Job#source_cluster_id}.
	SourceClusterId *string `field:"required" json:"sourceClusterId" yaml:"sourceClusterId"`
}

