// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoInitScripts struct {
	// abfss block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/data-sources/cluster#abfss DataDatabricksCluster#abfss}
	Abfss *DataDatabricksClusterClusterInfoInitScriptsAbfss `field:"optional" json:"abfss" yaml:"abfss"`
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/data-sources/cluster#dbfs DataDatabricksCluster#dbfs}
	Dbfs *DataDatabricksClusterClusterInfoInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/data-sources/cluster#file DataDatabricksCluster#file}
	File *DataDatabricksClusterClusterInfoInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/data-sources/cluster#gcs DataDatabricksCluster#gcs}
	Gcs *DataDatabricksClusterClusterInfoInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/data-sources/cluster#s3 DataDatabricksCluster#s3}
	S3 *DataDatabricksClusterClusterInfoInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/data-sources/cluster#volumes DataDatabricksCluster#volumes}
	Volumes *DataDatabricksClusterClusterInfoInitScriptsVolumes `field:"optional" json:"volumes" yaml:"volumes"`
	// workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.33.0/docs/data-sources/cluster#workspace DataDatabricksCluster#workspace}
	Workspace *DataDatabricksClusterClusterInfoInitScriptsWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}

