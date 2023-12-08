// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.1/docs/data-sources/job#created_time DataDatabricksJob#created_time}.
	CreatedTime *float64 `field:"optional" json:"createdTime" yaml:"createdTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.1/docs/data-sources/job#creator_user_name DataDatabricksJob#creator_user_name}.
	CreatorUserName *string `field:"optional" json:"creatorUserName" yaml:"creatorUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.1/docs/data-sources/job#job_id DataDatabricksJob#job_id}.
	JobId *float64 `field:"optional" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.1/docs/data-sources/job#run_as_user_name DataDatabricksJob#run_as_user_name}.
	RunAsUserName *string `field:"optional" json:"runAsUserName" yaml:"runAsUserName"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.1/docs/data-sources/job#settings DataDatabricksJob#settings}
	Settings *DataDatabricksJobJobSettingsSettings `field:"optional" json:"settings" yaml:"settings"`
}

