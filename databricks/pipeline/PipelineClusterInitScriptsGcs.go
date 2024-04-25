// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineClusterInitScriptsGcs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.41.0/docs/resources/pipeline#destination Pipeline#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

