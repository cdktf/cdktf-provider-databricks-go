// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineClusterAutoscale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/pipeline#max_workers Pipeline#max_workers}.
	MaxWorkers *float64 `field:"required" json:"maxWorkers" yaml:"maxWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/pipeline#min_workers Pipeline#min_workers}.
	MinWorkers *float64 `field:"required" json:"minWorkers" yaml:"minWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/pipeline#mode Pipeline#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

