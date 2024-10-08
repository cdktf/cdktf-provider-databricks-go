// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoSpecClusterLogConf struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/data-sources/cluster_pluginframework#dbfs DataDatabricksClusterPluginframework#dbfs}.
	Dbfs *DataDatabricksClusterPluginframeworkClusterInfoSpecClusterLogConfDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/data-sources/cluster_pluginframework#s3 DataDatabricksClusterPluginframework#s3}.
	S3 *DataDatabricksClusterPluginframeworkClusterInfoSpecClusterLogConfS3 `field:"optional" json:"s3" yaml:"s3"`
}

