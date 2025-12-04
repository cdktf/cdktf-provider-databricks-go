// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobGitSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#url Job#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#branch Job#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#commit Job#commit}.
	Commit *string `field:"optional" json:"commit" yaml:"commit"`
	// git_snapshot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#git_snapshot Job#git_snapshot}
	GitSnapshot *JobGitSourceGitSnapshot `field:"optional" json:"gitSnapshot" yaml:"gitSnapshot"`
	// job_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#job_source Job#job_source}
	JobSource *JobGitSourceJobSource `field:"optional" json:"jobSource" yaml:"jobSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#provider Job#provider}.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#tag Job#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

