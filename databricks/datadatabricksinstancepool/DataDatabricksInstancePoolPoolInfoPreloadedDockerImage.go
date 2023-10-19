// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoPreloadedDockerImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.28.1/docs/data-sources/instance_pool#url DataDatabricksInstancePool#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.28.1/docs/data-sources/instance_pool#basic_auth DataDatabricksInstancePool#basic_auth}
	BasicAuth *DataDatabricksInstancePoolPoolInfoPreloadedDockerImageBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

