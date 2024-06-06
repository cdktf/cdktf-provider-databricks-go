// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobGitSourceGitSnapshot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/resources/job#used_commit Job#used_commit}.
	UsedCommit *string `field:"optional" json:"usedCommit" yaml:"usedCommit"`
}

