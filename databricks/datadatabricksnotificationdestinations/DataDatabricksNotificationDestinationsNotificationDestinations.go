// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksnotificationdestinations


type DataDatabricksNotificationDestinationsNotificationDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/notification_destinations#destination_type DataDatabricksNotificationDestinations#destination_type}.
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/notification_destinations#display_name DataDatabricksNotificationDestinations#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/notification_destinations#id DataDatabricksNotificationDestinations#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

