// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskNotebookTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/data-sources/job#notebook_path DataDatabricksJob#notebook_path}.
	NotebookPath *string `field:"required" json:"notebookPath" yaml:"notebookPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/data-sources/job#base_parameters DataDatabricksJob#base_parameters}.
	BaseParameters *map[string]*string `field:"optional" json:"baseParameters" yaml:"baseParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/data-sources/job#source DataDatabricksJob#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

