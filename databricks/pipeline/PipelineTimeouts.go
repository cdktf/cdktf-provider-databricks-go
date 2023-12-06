// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/pipeline#default Pipeline#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
}

