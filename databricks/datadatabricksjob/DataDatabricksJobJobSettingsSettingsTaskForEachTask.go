// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/job#inputs DataDatabricksJob#inputs}.
	Inputs *string `field:"required" json:"inputs" yaml:"inputs"`
	// task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/job#task DataDatabricksJob#task}
	Task *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTask `field:"required" json:"task" yaml:"task"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/data-sources/job#concurrency DataDatabricksJob#concurrency}.
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
}

