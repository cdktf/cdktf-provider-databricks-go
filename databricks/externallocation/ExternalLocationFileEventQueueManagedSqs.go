// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externallocation


type ExternalLocationFileEventQueueManagedSqs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/external_location#queue_url ExternalLocation#queue_url}.
	QueueUrl *string `field:"optional" json:"queueUrl" yaml:"queueUrl"`
}

