// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repo


type RepoSparseCheckout struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/repo#patterns Repo#patterns}.
	Patterns *[]*string `field:"required" json:"patterns" yaml:"patterns"`
}

