// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterDockerImageBasicAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/cluster#password Cluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/cluster#username Cluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

