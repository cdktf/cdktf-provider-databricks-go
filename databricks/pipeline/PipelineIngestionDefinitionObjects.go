// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjects struct {
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/pipeline#schema Pipeline#schema}
	Schema *PipelineIngestionDefinitionObjectsSchema `field:"optional" json:"schema" yaml:"schema"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/pipeline#table Pipeline#table}
	Table *PipelineIngestionDefinitionObjectsTable `field:"optional" json:"table" yaml:"table"`
}

