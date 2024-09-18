// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoInitScripts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#abfss DataDatabricksClusterPluginframework#abfss}.
	Abfss *DataDatabricksClusterPluginframeworkClusterInfoInitScriptsAbfss `field:"optional" json:"abfss" yaml:"abfss"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#dbfs DataDatabricksClusterPluginframework#dbfs}.
	Dbfs *DataDatabricksClusterPluginframeworkClusterInfoInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#file DataDatabricksClusterPluginframework#file}.
	File *DataDatabricksClusterPluginframeworkClusterInfoInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#gcs DataDatabricksClusterPluginframework#gcs}.
	Gcs *DataDatabricksClusterPluginframeworkClusterInfoInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#s3 DataDatabricksClusterPluginframework#s3}.
	S3 *DataDatabricksClusterPluginframeworkClusterInfoInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#volumes DataDatabricksClusterPluginframework#volumes}.
	Volumes *DataDatabricksClusterPluginframeworkClusterInfoInitScriptsVolumes `field:"optional" json:"volumes" yaml:"volumes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#workspace DataDatabricksClusterPluginframework#workspace}.
	Workspace *DataDatabricksClusterPluginframeworkClusterInfoInitScriptsWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}

