// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskNewClusterInitScripts struct {
	// abfss block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#abfss Job#abfss}
	Abfss *JobTaskForEachTaskTaskNewClusterInitScriptsAbfss `field:"optional" json:"abfss" yaml:"abfss"`
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#dbfs Job#dbfs}
	Dbfs *JobTaskForEachTaskTaskNewClusterInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#file Job#file}
	File *JobTaskForEachTaskTaskNewClusterInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#gcs Job#gcs}
	Gcs *JobTaskForEachTaskTaskNewClusterInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#s3 Job#s3}
	S3 *JobTaskForEachTaskTaskNewClusterInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#volumes Job#volumes}
	Volumes *JobTaskForEachTaskTaskNewClusterInitScriptsVolumes `field:"optional" json:"volumes" yaml:"volumes"`
	// workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#workspace Job#workspace}
	Workspace *JobTaskForEachTaskTaskNewClusterInitScriptsWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}

