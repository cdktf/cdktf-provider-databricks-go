// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineDeployment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/pipeline#kind Pipeline#kind}.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/pipeline#metadata_file_path Pipeline#metadata_file_path}.
	MetadataFilePath *string `field:"optional" json:"metadataFilePath" yaml:"metadataFilePath"`
}

