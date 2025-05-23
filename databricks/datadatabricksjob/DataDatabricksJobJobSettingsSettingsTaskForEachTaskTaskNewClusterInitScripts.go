// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScripts struct {
	// abfss block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#abfss DataDatabricksJob#abfss}
	Abfss *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsAbfss `field:"optional" json:"abfss" yaml:"abfss"`
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#dbfs DataDatabricksJob#dbfs}
	Dbfs *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#file DataDatabricksJob#file}
	File *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#gcs DataDatabricksJob#gcs}
	Gcs *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#s3 DataDatabricksJob#s3}
	S3 *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#volumes DataDatabricksJob#volumes}
	Volumes *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsVolumes `field:"optional" json:"volumes" yaml:"volumes"`
	// workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/job#workspace DataDatabricksJob#workspace}
	Workspace *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}

