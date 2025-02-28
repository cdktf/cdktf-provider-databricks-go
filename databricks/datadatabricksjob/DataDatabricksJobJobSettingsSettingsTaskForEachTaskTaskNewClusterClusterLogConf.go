// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterClusterLogConf struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/job#dbfs DataDatabricksJob#dbfs}
	Dbfs *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterClusterLogConfDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/job#s3 DataDatabricksJob#s3}
	S3 *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterClusterLogConfS3 `field:"optional" json:"s3" yaml:"s3"`
}

