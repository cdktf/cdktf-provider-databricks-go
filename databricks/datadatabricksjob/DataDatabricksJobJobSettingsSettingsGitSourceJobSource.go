// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsGitSourceJobSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/data-sources/job#import_from_git_branch DataDatabricksJob#import_from_git_branch}.
	ImportFromGitBranch *string `field:"required" json:"importFromGitBranch" yaml:"importFromGitBranch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/data-sources/job#job_config_path DataDatabricksJob#job_config_path}.
	JobConfigPath *string `field:"required" json:"jobConfigPath" yaml:"jobConfigPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/data-sources/job#dirty_state DataDatabricksJob#dirty_state}.
	DirtyState *string `field:"optional" json:"dirtyState" yaml:"dirtyState"`
}

