// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationdestination


type NotificationDestinationConfigSlack struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/notification_destination#channel_id NotificationDestination#channel_id}.
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/notification_destination#channel_id_set NotificationDestination#channel_id_set}.
	ChannelIdSet interface{} `field:"optional" json:"channelIdSet" yaml:"channelIdSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/notification_destination#oauth_token NotificationDestination#oauth_token}.
	OauthToken *string `field:"optional" json:"oauthToken" yaml:"oauthToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/notification_destination#oauth_token_set NotificationDestination#oauth_token_set}.
	OauthTokenSet interface{} `field:"optional" json:"oauthTokenSet" yaml:"oauthTokenSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/notification_destination#url NotificationDestination#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/notification_destination#url_set NotificationDestination#url_set}.
	UrlSet interface{} `field:"optional" json:"urlSet" yaml:"urlSet"`
}

