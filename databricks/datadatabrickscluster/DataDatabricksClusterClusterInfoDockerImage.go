// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoDockerImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/data-sources/cluster#url DataDatabricksCluster#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/data-sources/cluster#basic_auth DataDatabricksCluster#basic_auth}
	BasicAuth *DataDatabricksClusterClusterInfoDockerImageBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

