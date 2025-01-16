// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsworkspaces

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsWorkspacesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#account_id MwsWorkspaces#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#workspace_name MwsWorkspaces#workspace_name}.
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#aws_region MwsWorkspaces#aws_region}.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#cloud MwsWorkspaces#cloud}.
	Cloud *string `field:"optional" json:"cloud" yaml:"cloud"`
	// cloud_resource_container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#cloud_resource_container MwsWorkspaces#cloud_resource_container}
	CloudResourceContainer *MwsWorkspacesCloudResourceContainer `field:"optional" json:"cloudResourceContainer" yaml:"cloudResourceContainer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#creation_time MwsWorkspaces#creation_time}.
	CreationTime *float64 `field:"optional" json:"creationTime" yaml:"creationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#credentials_id MwsWorkspaces#credentials_id}.
	CredentialsId *string `field:"optional" json:"credentialsId" yaml:"credentialsId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#customer_managed_key_id MwsWorkspaces#customer_managed_key_id}.
	CustomerManagedKeyId *string `field:"optional" json:"customerManagedKeyId" yaml:"customerManagedKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#custom_tags MwsWorkspaces#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#deployment_name MwsWorkspaces#deployment_name}.
	DeploymentName *string `field:"optional" json:"deploymentName" yaml:"deploymentName"`
	// external_customer_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#external_customer_info MwsWorkspaces#external_customer_info}
	ExternalCustomerInfo *MwsWorkspacesExternalCustomerInfo `field:"optional" json:"externalCustomerInfo" yaml:"externalCustomerInfo"`
	// gcp_managed_network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#gcp_managed_network_config MwsWorkspaces#gcp_managed_network_config}
	GcpManagedNetworkConfig *MwsWorkspacesGcpManagedNetworkConfig `field:"optional" json:"gcpManagedNetworkConfig" yaml:"gcpManagedNetworkConfig"`
	// gke_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#gke_config MwsWorkspaces#gke_config}
	GkeConfig *MwsWorkspacesGkeConfig `field:"optional" json:"gkeConfig" yaml:"gkeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#id MwsWorkspaces#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#is_no_public_ip_enabled MwsWorkspaces#is_no_public_ip_enabled}.
	IsNoPublicIpEnabled interface{} `field:"optional" json:"isNoPublicIpEnabled" yaml:"isNoPublicIpEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#location MwsWorkspaces#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#managed_services_customer_managed_key_id MwsWorkspaces#managed_services_customer_managed_key_id}.
	ManagedServicesCustomerManagedKeyId *string `field:"optional" json:"managedServicesCustomerManagedKeyId" yaml:"managedServicesCustomerManagedKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#network_id MwsWorkspaces#network_id}.
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#pricing_tier MwsWorkspaces#pricing_tier}.
	PricingTier *string `field:"optional" json:"pricingTier" yaml:"pricingTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#private_access_settings_id MwsWorkspaces#private_access_settings_id}.
	PrivateAccessSettingsId *string `field:"optional" json:"privateAccessSettingsId" yaml:"privateAccessSettingsId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#storage_configuration_id MwsWorkspaces#storage_configuration_id}.
	StorageConfigurationId *string `field:"optional" json:"storageConfigurationId" yaml:"storageConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#storage_customer_managed_key_id MwsWorkspaces#storage_customer_managed_key_id}.
	StorageCustomerManagedKeyId *string `field:"optional" json:"storageCustomerManagedKeyId" yaml:"storageCustomerManagedKeyId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#timeouts MwsWorkspaces#timeouts}
	Timeouts *MwsWorkspacesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#token MwsWorkspaces#token}
	Token *MwsWorkspacesToken `field:"optional" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#workspace_id MwsWorkspaces#workspace_id}.
	WorkspaceId *float64 `field:"optional" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#workspace_status MwsWorkspaces#workspace_status}.
	WorkspaceStatus *string `field:"optional" json:"workspaceStatus" yaml:"workspaceStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#workspace_status_message MwsWorkspaces#workspace_status_message}.
	WorkspaceStatusMessage *string `field:"optional" json:"workspaceStatusMessage" yaml:"workspaceStatusMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.63.0/docs/resources/mws_workspaces#workspace_url MwsWorkspaces#workspace_url}.
	WorkspaceUrl *string `field:"optional" json:"workspaceUrl" yaml:"workspaceUrl"`
}

