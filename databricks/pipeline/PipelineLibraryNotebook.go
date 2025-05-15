// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineLibraryNotebook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/resources/pipeline#path Pipeline#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

