// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskPipelineTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/data-sources/job#pipeline_id DataDatabricksJob#pipeline_id}.
	PipelineId *string `field:"required" json:"pipelineId" yaml:"pipelineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/data-sources/job#full_refresh DataDatabricksJob#full_refresh}.
	FullRefresh interface{} `field:"optional" json:"fullRefresh" yaml:"fullRefresh"`
}

