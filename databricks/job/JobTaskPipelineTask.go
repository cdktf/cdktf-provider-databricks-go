// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskPipelineTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/resources/job#pipeline_id Job#pipeline_id}.
	PipelineId *string `field:"required" json:"pipelineId" yaml:"pipelineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/resources/job#full_refresh Job#full_refresh}.
	FullRefresh interface{} `field:"optional" json:"fullRefresh" yaml:"fullRefresh"`
}

