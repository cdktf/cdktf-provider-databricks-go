// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoPreloadedDockerImageBasicAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/instance_pool#password DataDatabricksInstancePool#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/instance_pool#username DataDatabricksInstancePool#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

