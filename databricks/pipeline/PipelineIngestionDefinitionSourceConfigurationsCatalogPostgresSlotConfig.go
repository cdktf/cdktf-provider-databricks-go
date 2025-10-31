// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionSourceConfigurationsCatalogPostgresSlotConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/pipeline#publication_name Pipeline#publication_name}.
	PublicationName *string `field:"optional" json:"publicationName" yaml:"publicationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/pipeline#slot_name Pipeline#slot_name}.
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
}

