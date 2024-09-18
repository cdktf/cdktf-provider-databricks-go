// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoSpecInitScripts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#abfss DataDatabricksClusterPluginframework#abfss}.
	Abfss *DataDatabricksClusterPluginframeworkClusterInfoSpecInitScriptsAbfss `field:"optional" json:"abfss" yaml:"abfss"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#dbfs DataDatabricksClusterPluginframework#dbfs}.
	Dbfs *DataDatabricksClusterPluginframeworkClusterInfoSpecInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#file DataDatabricksClusterPluginframework#file}.
	File *DataDatabricksClusterPluginframeworkClusterInfoSpecInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#gcs DataDatabricksClusterPluginframework#gcs}.
	Gcs *DataDatabricksClusterPluginframeworkClusterInfoSpecInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#s3 DataDatabricksClusterPluginframework#s3}.
	S3 *DataDatabricksClusterPluginframeworkClusterInfoSpecInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#volumes DataDatabricksClusterPluginframework#volumes}.
	Volumes *DataDatabricksClusterPluginframeworkClusterInfoSpecInitScriptsVolumes `field:"optional" json:"volumes" yaml:"volumes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/cluster_pluginframework#workspace DataDatabricksClusterPluginframework#workspace}.
	Workspace *DataDatabricksClusterPluginframeworkClusterInfoSpecInitScriptsWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}

