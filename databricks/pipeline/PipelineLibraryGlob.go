// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineLibraryGlob struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/resources/pipeline#include Pipeline#include}.
	Include *string `field:"optional" json:"include" yaml:"include"`
}

