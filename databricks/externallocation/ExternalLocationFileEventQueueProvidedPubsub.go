// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externallocation


type ExternalLocationFileEventQueueProvidedPubsub struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/external_location#subscription_name ExternalLocation#subscription_name}.
	SubscriptionName *string `field:"required" json:"subscriptionName" yaml:"subscriptionName"`
}

