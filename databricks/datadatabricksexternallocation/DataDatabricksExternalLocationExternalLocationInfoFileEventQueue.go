// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksexternallocation


type DataDatabricksExternalLocationExternalLocationInfoFileEventQueue struct {
	// managed_aqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/external_location#managed_aqs DataDatabricksExternalLocation#managed_aqs}
	ManagedAqs *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedAqs `field:"optional" json:"managedAqs" yaml:"managedAqs"`
	// managed_pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/external_location#managed_pubsub DataDatabricksExternalLocation#managed_pubsub}
	ManagedPubsub *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedPubsub `field:"optional" json:"managedPubsub" yaml:"managedPubsub"`
	// managed_sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/external_location#managed_sqs DataDatabricksExternalLocation#managed_sqs}
	ManagedSqs *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueManagedSqs `field:"optional" json:"managedSqs" yaml:"managedSqs"`
	// provided_aqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/external_location#provided_aqs DataDatabricksExternalLocation#provided_aqs}
	ProvidedAqs *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedAqs `field:"optional" json:"providedAqs" yaml:"providedAqs"`
	// provided_pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/external_location#provided_pubsub DataDatabricksExternalLocation#provided_pubsub}
	ProvidedPubsub *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedPubsub `field:"optional" json:"providedPubsub" yaml:"providedPubsub"`
	// provided_sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/external_location#provided_sqs DataDatabricksExternalLocation#provided_sqs}
	ProvidedSqs *DataDatabricksExternalLocationExternalLocationInfoFileEventQueueProvidedSqs `field:"optional" json:"providedSqs" yaml:"providedSqs"`
}

