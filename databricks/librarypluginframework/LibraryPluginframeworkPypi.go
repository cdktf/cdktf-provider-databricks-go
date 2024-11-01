// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package librarypluginframework


type LibraryPluginframeworkPypi struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.56.0/docs/resources/library_pluginframework#package LibraryPluginframework#package}.
	Package *string `field:"required" json:"package" yaml:"package"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.56.0/docs/resources/library_pluginframework#repo LibraryPluginframework#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

