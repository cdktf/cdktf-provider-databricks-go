// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationdestination


type NotificationDestinationConfigPagerduty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/notification_destination#integration_key NotificationDestination#integration_key}.
	IntegrationKey *string `field:"optional" json:"integrationKey" yaml:"integrationKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.51.0/docs/resources/notification_destination#integration_key_set NotificationDestination#integration_key_set}.
	IntegrationKeySet interface{} `field:"optional" json:"integrationKeySet" yaml:"integrationKeySet"`
}

