// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionSourceConfigurationsCatalogPostgres struct {
	// slot_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/pipeline#slot_config Pipeline#slot_config}
	SlotConfig *PipelineIngestionDefinitionSourceConfigurationsCatalogPostgresSlotConfig `field:"optional" json:"slotConfig" yaml:"slotConfig"`
}

