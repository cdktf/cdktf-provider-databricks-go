// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionSourceConfigurations struct {
	// catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/pipeline#catalog Pipeline#catalog}
	Catalog *PipelineIngestionDefinitionSourceConfigurationsCatalog `field:"optional" json:"catalog" yaml:"catalog"`
}

