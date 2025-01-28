// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjects struct {
	// report block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.0/docs/resources/pipeline#report Pipeline#report}
	Report *PipelineIngestionDefinitionObjectsReport `field:"optional" json:"report" yaml:"report"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.0/docs/resources/pipeline#schema Pipeline#schema}
	Schema *PipelineIngestionDefinitionObjectsSchema `field:"optional" json:"schema" yaml:"schema"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.0/docs/resources/pipeline#table Pipeline#table}
	Table *PipelineIngestionDefinitionObjectsTable `field:"optional" json:"table" yaml:"table"`
}

