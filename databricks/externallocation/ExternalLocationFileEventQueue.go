// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externallocation


type ExternalLocationFileEventQueue struct {
	// managed_aqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/external_location#managed_aqs ExternalLocation#managed_aqs}
	ManagedAqs *ExternalLocationFileEventQueueManagedAqs `field:"optional" json:"managedAqs" yaml:"managedAqs"`
	// managed_pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/external_location#managed_pubsub ExternalLocation#managed_pubsub}
	ManagedPubsub *ExternalLocationFileEventQueueManagedPubsub `field:"optional" json:"managedPubsub" yaml:"managedPubsub"`
	// managed_sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/external_location#managed_sqs ExternalLocation#managed_sqs}
	ManagedSqs *ExternalLocationFileEventQueueManagedSqs `field:"optional" json:"managedSqs" yaml:"managedSqs"`
	// provided_aqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/external_location#provided_aqs ExternalLocation#provided_aqs}
	ProvidedAqs *ExternalLocationFileEventQueueProvidedAqs `field:"optional" json:"providedAqs" yaml:"providedAqs"`
	// provided_pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/external_location#provided_pubsub ExternalLocation#provided_pubsub}
	ProvidedPubsub *ExternalLocationFileEventQueueProvidedPubsub `field:"optional" json:"providedPubsub" yaml:"providedPubsub"`
	// provided_sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/external_location#provided_sqs ExternalLocation#provided_sqs}
	ProvidedSqs *ExternalLocationFileEventQueueProvidedSqs `field:"optional" json:"providedSqs" yaml:"providedSqs"`
}

