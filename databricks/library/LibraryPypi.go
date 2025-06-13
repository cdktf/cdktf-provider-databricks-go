// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package library


type LibraryPypi struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/library#package Library#package}.
	Package *string `field:"required" json:"package" yaml:"package"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/library#repo Library#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

