// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationdestination


type NotificationDestinationConfigMicrosoftTeams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/notification_destination#app_id NotificationDestination#app_id}.
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/notification_destination#app_id_set NotificationDestination#app_id_set}.
	AppIdSet interface{} `field:"optional" json:"appIdSet" yaml:"appIdSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/notification_destination#auth_secret NotificationDestination#auth_secret}.
	AuthSecret *string `field:"optional" json:"authSecret" yaml:"authSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/notification_destination#auth_secret_set NotificationDestination#auth_secret_set}.
	AuthSecretSet interface{} `field:"optional" json:"authSecretSet" yaml:"authSecretSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/notification_destination#channel_url NotificationDestination#channel_url}.
	ChannelUrl *string `field:"optional" json:"channelUrl" yaml:"channelUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/notification_destination#channel_url_set NotificationDestination#channel_url_set}.
	ChannelUrlSet interface{} `field:"optional" json:"channelUrlSet" yaml:"channelUrlSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/notification_destination#tenant_id NotificationDestination#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/notification_destination#tenant_id_set NotificationDestination#tenant_id_set}.
	TenantIdSet interface{} `field:"optional" json:"tenantIdSet" yaml:"tenantIdSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/notification_destination#url NotificationDestination#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/notification_destination#url_set NotificationDestination#url_set}.
	UrlSet interface{} `field:"optional" json:"urlSet" yaml:"urlSet"`
}

