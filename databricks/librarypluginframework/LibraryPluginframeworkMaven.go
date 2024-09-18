// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package librarypluginframework


type LibraryPluginframeworkMaven struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/library_pluginframework#coordinates LibraryPluginframework#coordinates}.
	Coordinates *string `field:"required" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/library_pluginframework#exclusions LibraryPluginframework#exclusions}.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/library_pluginframework#repo LibraryPluginframework#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

