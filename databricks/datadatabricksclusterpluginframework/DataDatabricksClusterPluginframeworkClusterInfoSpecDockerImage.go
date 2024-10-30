// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoSpecDockerImage struct {
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.56.0/docs/data-sources/cluster_pluginframework#basic_auth DataDatabricksClusterPluginframework#basic_auth}
	BasicAuth interface{} `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.56.0/docs/data-sources/cluster_pluginframework#url DataDatabricksClusterPluginframework#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

