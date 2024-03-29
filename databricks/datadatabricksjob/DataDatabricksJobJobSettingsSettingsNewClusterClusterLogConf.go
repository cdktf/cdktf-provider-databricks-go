// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsNewClusterClusterLogConf struct {
	// dbfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/data-sources/job#dbfs DataDatabricksJob#dbfs}
	Dbfs *DataDatabricksJobJobSettingsSettingsNewClusterClusterLogConfDbfs `field:"optional" json:"dbfs" yaml:"dbfs"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/data-sources/job#s3 DataDatabricksJob#s3}
	S3 *DataDatabricksJobJobSettingsSettingsNewClusterClusterLogConfS3 `field:"optional" json:"s3" yaml:"s3"`
}

