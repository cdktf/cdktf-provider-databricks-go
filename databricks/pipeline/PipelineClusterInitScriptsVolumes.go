// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineClusterInitScriptsVolumes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.28.0/docs/resources/pipeline#destination Pipeline#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}
