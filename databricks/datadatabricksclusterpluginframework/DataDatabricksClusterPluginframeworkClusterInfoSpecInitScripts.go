// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoSpecInitScripts struct {
	// abfss block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster_pluginframework#abfss DataDatabricksClusterPluginframework#abfss}
	Abfss interface{} `field:"optional" json:"abfss" yaml:"abfss"`
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster_pluginframework#dbfs DataDatabricksClusterPluginframework#dbfs}
	Dbfs interface{} `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster_pluginframework#file DataDatabricksClusterPluginframework#file}
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster_pluginframework#gcs DataDatabricksClusterPluginframework#gcs}
	Gcs interface{} `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster_pluginframework#s3 DataDatabricksClusterPluginframework#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster_pluginframework#volumes DataDatabricksClusterPluginframework#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
	// workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster_pluginframework#workspace DataDatabricksClusterPluginframework#workspace}
	Workspace interface{} `field:"optional" json:"workspace" yaml:"workspace"`
}

