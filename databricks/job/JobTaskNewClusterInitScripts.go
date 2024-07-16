// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskNewClusterInitScripts struct {
	// abfss block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/job#abfss Job#abfss}
	Abfss *JobTaskNewClusterInitScriptsAbfss `field:"optional" json:"abfss" yaml:"abfss"`
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/job#dbfs Job#dbfs}
	Dbfs *JobTaskNewClusterInitScriptsDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/job#file Job#file}
	File *JobTaskNewClusterInitScriptsFile `field:"optional" json:"file" yaml:"file"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/job#gcs Job#gcs}
	Gcs *JobTaskNewClusterInitScriptsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/job#s3 Job#s3}
	S3 *JobTaskNewClusterInitScriptsS3 `field:"optional" json:"s3" yaml:"s3"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/job#volumes Job#volumes}
	Volumes *JobTaskNewClusterInitScriptsVolumes `field:"optional" json:"volumes" yaml:"volumes"`
	// workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/job#workspace Job#workspace}
	Workspace *JobTaskNewClusterInitScriptsWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}

