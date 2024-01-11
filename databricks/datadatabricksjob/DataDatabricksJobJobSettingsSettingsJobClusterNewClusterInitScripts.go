// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScripts struct {
	// abfss block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/data-sources/job#abfss DataDatabricksJob#abfss}
	Abfss *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsAbfss `field:"optional" json:"abfss" yaml:"abfss"`
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/data-sources/job#dbfs DataDatabricksJob#dbfs}
	Dbfs *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/data-sources/job#file DataDatabricksJob#file}
	File *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/data-sources/job#gcs DataDatabricksJob#gcs}
	Gcs *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/data-sources/job#s3 DataDatabricksJob#s3}
	S3 *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/data-sources/job#volumes DataDatabricksJob#volumes}
	Volumes *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsVolumes `field:"optional" json:"volumes" yaml:"volumes"`
	// workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/data-sources/job#workspace DataDatabricksJob#workspace}
	Workspace *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}

