// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsReport struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.69.0/docs/resources/pipeline#destination_catalog Pipeline#destination_catalog}.
	DestinationCatalog *string `field:"optional" json:"destinationCatalog" yaml:"destinationCatalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.69.0/docs/resources/pipeline#destination_schema Pipeline#destination_schema}.
	DestinationSchema *string `field:"optional" json:"destinationSchema" yaml:"destinationSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.69.0/docs/resources/pipeline#destination_table Pipeline#destination_table}.
	DestinationTable *string `field:"optional" json:"destinationTable" yaml:"destinationTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.69.0/docs/resources/pipeline#source_url Pipeline#source_url}.
	SourceUrl *string `field:"optional" json:"sourceUrl" yaml:"sourceUrl"`
	// table_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.69.0/docs/resources/pipeline#table_configuration Pipeline#table_configuration}
	TableConfiguration *PipelineIngestionDefinitionObjectsReportTableConfiguration `field:"optional" json:"tableConfiguration" yaml:"tableConfiguration"`
}

