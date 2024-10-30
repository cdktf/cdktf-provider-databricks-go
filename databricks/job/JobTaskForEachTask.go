// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.56.0/docs/resources/job#inputs Job#inputs}.
	Inputs *string `field:"required" json:"inputs" yaml:"inputs"`
	// task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.56.0/docs/resources/job#task Job#task}
	Task *JobTaskForEachTaskTask `field:"required" json:"task" yaml:"task"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.56.0/docs/resources/job#concurrency Job#concurrency}.
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
}

