// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoSpecClusterLogConf struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster_pluginframework#dbfs DataDatabricksClusterPluginframework#dbfs}
	Dbfs interface{} `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster_pluginframework#s3 DataDatabricksClusterPluginframework#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

