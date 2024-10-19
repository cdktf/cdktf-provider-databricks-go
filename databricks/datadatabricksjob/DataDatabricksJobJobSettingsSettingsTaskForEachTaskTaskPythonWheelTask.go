// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPythonWheelTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.54.0/docs/data-sources/job#entry_point DataDatabricksJob#entry_point}.
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.54.0/docs/data-sources/job#named_parameters DataDatabricksJob#named_parameters}.
	NamedParameters *map[string]*string `field:"optional" json:"namedParameters" yaml:"namedParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.54.0/docs/data-sources/job#package_name DataDatabricksJob#package_name}.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.54.0/docs/data-sources/job#parameters DataDatabricksJob#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

