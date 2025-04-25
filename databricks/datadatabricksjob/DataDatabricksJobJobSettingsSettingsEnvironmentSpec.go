// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsEnvironmentSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/job#client DataDatabricksJob#client}.
	Client *string `field:"required" json:"client" yaml:"client"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/job#dependencies DataDatabricksJob#dependencies}.
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/job#jar_dependencies DataDatabricksJob#jar_dependencies}.
	JarDependencies *[]*string `field:"optional" json:"jarDependencies" yaml:"jarDependencies"`
}

