// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline#destination_catalog Pipeline#destination_catalog}.
	DestinationCatalog *string `field:"required" json:"destinationCatalog" yaml:"destinationCatalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline#destination_schema Pipeline#destination_schema}.
	DestinationSchema *string `field:"required" json:"destinationSchema" yaml:"destinationSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline#source_table Pipeline#source_table}.
	SourceTable *string `field:"required" json:"sourceTable" yaml:"sourceTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline#destination_table Pipeline#destination_table}.
	DestinationTable *string `field:"optional" json:"destinationTable" yaml:"destinationTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline#source_catalog Pipeline#source_catalog}.
	SourceCatalog *string `field:"optional" json:"sourceCatalog" yaml:"sourceCatalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline#source_schema Pipeline#source_schema}.
	SourceSchema *string `field:"optional" json:"sourceSchema" yaml:"sourceSchema"`
	// table_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline#table_configuration Pipeline#table_configuration}
	TableConfiguration *PipelineIngestionDefinitionObjectsTableTableConfiguration `field:"optional" json:"tableConfiguration" yaml:"tableConfiguration"`
}

