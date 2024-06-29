// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDependsOn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.2/docs/data-sources/job#task_key DataDatabricksJob#task_key}.
	TaskKey *string `field:"required" json:"taskKey" yaml:"taskKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.2/docs/data-sources/job#outcome DataDatabricksJob#outcome}.
	Outcome *string `field:"optional" json:"outcome" yaml:"outcome"`
}

