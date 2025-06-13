// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/pipeline#destination_catalog Pipeline#destination_catalog}.
	DestinationCatalog *string `field:"required" json:"destinationCatalog" yaml:"destinationCatalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/pipeline#destination_schema Pipeline#destination_schema}.
	DestinationSchema *string `field:"required" json:"destinationSchema" yaml:"destinationSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/pipeline#source_schema Pipeline#source_schema}.
	SourceSchema *string `field:"required" json:"sourceSchema" yaml:"sourceSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/pipeline#source_catalog Pipeline#source_catalog}.
	SourceCatalog *string `field:"optional" json:"sourceCatalog" yaml:"sourceCatalog"`
	// table_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/pipeline#table_configuration Pipeline#table_configuration}
	TableConfiguration *PipelineIngestionDefinitionObjectsSchemaTableConfiguration `field:"optional" json:"tableConfiguration" yaml:"tableConfiguration"`
}

