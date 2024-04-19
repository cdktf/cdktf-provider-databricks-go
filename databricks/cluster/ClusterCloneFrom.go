// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterCloneFrom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.40.0/docs/resources/cluster#source_cluster_id Cluster#source_cluster_id}.
	SourceClusterId *string `field:"required" json:"sourceClusterId" yaml:"sourceClusterId"`
}

