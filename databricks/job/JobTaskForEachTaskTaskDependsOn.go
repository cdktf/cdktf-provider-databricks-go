// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskDependsOn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.72.0/docs/resources/job#task_key Job#task_key}.
	TaskKey *string `field:"required" json:"taskKey" yaml:"taskKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.72.0/docs/resources/job#outcome Job#outcome}.
	Outcome *string `field:"optional" json:"outcome" yaml:"outcome"`
}

