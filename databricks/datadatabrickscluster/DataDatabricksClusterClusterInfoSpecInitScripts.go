// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoSpecInitScripts struct {
	// abfss block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster#abfss DataDatabricksCluster#abfss}
	Abfss *DataDatabricksClusterClusterInfoSpecInitScriptsAbfss `field:"optional" json:"abfss" yaml:"abfss"`
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster#dbfs DataDatabricksCluster#dbfs}
	Dbfs *DataDatabricksClusterClusterInfoSpecInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster#file DataDatabricksCluster#file}
	File *DataDatabricksClusterClusterInfoSpecInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster#gcs DataDatabricksCluster#gcs}
	Gcs *DataDatabricksClusterClusterInfoSpecInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster#s3 DataDatabricksCluster#s3}
	S3 *DataDatabricksClusterClusterInfoSpecInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster#volumes DataDatabricksCluster#volumes}
	Volumes *DataDatabricksClusterClusterInfoSpecInitScriptsVolumes `field:"optional" json:"volumes" yaml:"volumes"`
	// workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/cluster#workspace DataDatabricksCluster#workspace}
	Workspace *DataDatabricksClusterClusterInfoSpecInitScriptsWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}

