// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobNewClusterClusterLogConf struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/resources/job#dbfs Job#dbfs}
	Dbfs *JobNewClusterClusterLogConfDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/resources/job#s3 Job#s3}
	S3 *JobNewClusterClusterLogConfS3 `field:"optional" json:"s3" yaml:"s3"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/resources/job#volumes Job#volumes}
	Volumes *JobNewClusterClusterLogConfVolumes `field:"optional" json:"volumes" yaml:"volumes"`
}

