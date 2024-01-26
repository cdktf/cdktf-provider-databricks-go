// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobGitSourceJobSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.35.0/docs/resources/job#import_from_git_branch Job#import_from_git_branch}.
	ImportFromGitBranch *string `field:"required" json:"importFromGitBranch" yaml:"importFromGitBranch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.35.0/docs/resources/job#job_config_path Job#job_config_path}.
	JobConfigPath *string `field:"required" json:"jobConfigPath" yaml:"jobConfigPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.35.0/docs/resources/job#dirty_state Job#dirty_state}.
	DirtyState *string `field:"optional" json:"dirtyState" yaml:"dirtyState"`
}

