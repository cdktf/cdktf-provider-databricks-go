// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusters


type DataDatabricksClustersFilterBy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/data-sources/clusters#cluster_sources DataDatabricksClusters#cluster_sources}.
	ClusterSources *[]*string `field:"optional" json:"clusterSources" yaml:"clusterSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/data-sources/clusters#cluster_states DataDatabricksClusters#cluster_states}.
	ClusterStates *[]*string `field:"optional" json:"clusterStates" yaml:"clusterStates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/data-sources/clusters#is_pinned DataDatabricksClusters#is_pinned}.
	IsPinned interface{} `field:"optional" json:"isPinned" yaml:"isPinned"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/data-sources/clusters#policy_id DataDatabricksClusters#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
}

