// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSparkJarTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.49.0/docs/data-sources/job#jar_uri DataDatabricksJob#jar_uri}.
	JarUri *string `field:"optional" json:"jarUri" yaml:"jarUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.49.0/docs/data-sources/job#main_class_name DataDatabricksJob#main_class_name}.
	MainClassName *string `field:"optional" json:"mainClassName" yaml:"mainClassName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.49.0/docs/data-sources/job#parameters DataDatabricksJob#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

