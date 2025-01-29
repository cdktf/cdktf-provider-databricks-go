// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusterpolicy


type ClusterPolicyLibrariesMaven struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/resources/cluster_policy#coordinates ClusterPolicy#coordinates}.
	Coordinates *string `field:"required" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/resources/cluster_policy#exclusions ClusterPolicy#exclusions}.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/resources/cluster_policy#repo ClusterPolicy#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

