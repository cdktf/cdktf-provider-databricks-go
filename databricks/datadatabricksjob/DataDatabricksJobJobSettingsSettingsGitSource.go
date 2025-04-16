// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsGitSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/job#url DataDatabricksJob#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/job#branch DataDatabricksJob#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/job#commit DataDatabricksJob#commit}.
	Commit *string `field:"optional" json:"commit" yaml:"commit"`
	// job_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/job#job_source DataDatabricksJob#job_source}
	JobSource *DataDatabricksJobJobSettingsSettingsGitSourceJobSource `field:"optional" json:"jobSource" yaml:"jobSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/job#provider DataDatabricksJob#provider}.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/job#tag DataDatabricksJob#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

