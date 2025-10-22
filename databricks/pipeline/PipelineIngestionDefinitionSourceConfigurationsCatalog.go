// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionSourceConfigurationsCatalog struct {
	// postgres block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/pipeline#postgres Pipeline#postgres}
	Postgres *PipelineIngestionDefinitionSourceConfigurationsCatalogPostgres `field:"optional" json:"postgres" yaml:"postgres"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/pipeline#source_catalog Pipeline#source_catalog}.
	SourceCatalog *string `field:"optional" json:"sourceCatalog" yaml:"sourceCatalog"`
}

