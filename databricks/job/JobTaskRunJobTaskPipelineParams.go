// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskRunJobTaskPipelineParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/resources/job#full_refresh Job#full_refresh}.
	FullRefresh interface{} `field:"optional" json:"fullRefresh" yaml:"fullRefresh"`
}

