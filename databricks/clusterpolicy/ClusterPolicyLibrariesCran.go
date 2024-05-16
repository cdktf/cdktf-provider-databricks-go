// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusterpolicy


type ClusterPolicyLibrariesCran struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/resources/cluster_policy#package ClusterPolicy#package}.
	Package *string `field:"required" json:"package" yaml:"package"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/resources/cluster_policy#repo ClusterPolicy#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

