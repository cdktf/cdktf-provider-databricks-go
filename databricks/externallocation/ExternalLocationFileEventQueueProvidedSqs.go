// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externallocation


type ExternalLocationFileEventQueueProvidedSqs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/external_location#managed_resource_id ExternalLocation#managed_resource_id}.
	ManagedResourceId *string `field:"optional" json:"managedResourceId" yaml:"managedResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/external_location#queue_url ExternalLocation#queue_url}.
	QueueUrl *string `field:"optional" json:"queueUrl" yaml:"queueUrl"`
}

