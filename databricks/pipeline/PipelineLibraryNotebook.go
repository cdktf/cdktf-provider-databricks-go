// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineLibraryNotebook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/pipeline#path Pipeline#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

