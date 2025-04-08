// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoSpecDockerImageBasicAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.72.0/docs/data-sources/cluster#password DataDatabricksCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.72.0/docs/data-sources/cluster#username DataDatabricksCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

