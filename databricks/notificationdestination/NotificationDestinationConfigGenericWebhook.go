// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationdestination


type NotificationDestinationConfigGenericWebhook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/notification_destination#password NotificationDestination#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/notification_destination#password_set NotificationDestination#password_set}.
	PasswordSet interface{} `field:"optional" json:"passwordSet" yaml:"passwordSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/notification_destination#url NotificationDestination#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/notification_destination#url_set NotificationDestination#url_set}.
	UrlSet interface{} `field:"optional" json:"urlSet" yaml:"urlSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/notification_destination#username NotificationDestination#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/notification_destination#username_set NotificationDestination#username_set}.
	UsernameSet interface{} `field:"optional" json:"usernameSet" yaml:"usernameSet"`
}

