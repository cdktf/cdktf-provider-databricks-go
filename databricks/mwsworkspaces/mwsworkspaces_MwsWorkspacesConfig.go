package mwsworkspaces

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsWorkspacesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#account_id MwsWorkspaces#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#workspace_name MwsWorkspaces#workspace_name}.
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#aws_region MwsWorkspaces#aws_region}.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#cloud MwsWorkspaces#cloud}.
	Cloud *string `field:"optional" json:"cloud" yaml:"cloud"`
	// cloud_resource_bucket block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#cloud_resource_bucket MwsWorkspaces#cloud_resource_bucket}
	CloudResourceBucket *MwsWorkspacesCloudResourceBucket `field:"optional" json:"cloudResourceBucket" yaml:"cloudResourceBucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#creation_time MwsWorkspaces#creation_time}.
	CreationTime *float64 `field:"optional" json:"creationTime" yaml:"creationTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#credentials_id MwsWorkspaces#credentials_id}.
	CredentialsId *string `field:"optional" json:"credentialsId" yaml:"credentialsId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#customer_managed_key_id MwsWorkspaces#customer_managed_key_id}.
	CustomerManagedKeyId *string `field:"optional" json:"customerManagedKeyId" yaml:"customerManagedKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#deployment_name MwsWorkspaces#deployment_name}.
	DeploymentName *string `field:"optional" json:"deploymentName" yaml:"deploymentName"`
	// external_customer_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#external_customer_info MwsWorkspaces#external_customer_info}
	ExternalCustomerInfo *MwsWorkspacesExternalCustomerInfo `field:"optional" json:"externalCustomerInfo" yaml:"externalCustomerInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#id MwsWorkspaces#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#is_no_public_ip_enabled MwsWorkspaces#is_no_public_ip_enabled}.
	IsNoPublicIpEnabled interface{} `field:"optional" json:"isNoPublicIpEnabled" yaml:"isNoPublicIpEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#location MwsWorkspaces#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#managed_services_customer_managed_key_id MwsWorkspaces#managed_services_customer_managed_key_id}.
	ManagedServicesCustomerManagedKeyId *string `field:"optional" json:"managedServicesCustomerManagedKeyId" yaml:"managedServicesCustomerManagedKeyId"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#network MwsWorkspaces#network}
	Network *MwsWorkspacesNetwork `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#network_id MwsWorkspaces#network_id}.
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#pricing_tier MwsWorkspaces#pricing_tier}.
	PricingTier *string `field:"optional" json:"pricingTier" yaml:"pricingTier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#private_access_settings_id MwsWorkspaces#private_access_settings_id}.
	PrivateAccessSettingsId *string `field:"optional" json:"privateAccessSettingsId" yaml:"privateAccessSettingsId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#storage_configuration_id MwsWorkspaces#storage_configuration_id}.
	StorageConfigurationId *string `field:"optional" json:"storageConfigurationId" yaml:"storageConfigurationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#storage_customer_managed_key_id MwsWorkspaces#storage_customer_managed_key_id}.
	StorageCustomerManagedKeyId *string `field:"optional" json:"storageCustomerManagedKeyId" yaml:"storageCustomerManagedKeyId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#timeouts MwsWorkspaces#timeouts}
	Timeouts *MwsWorkspacesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// token block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#token MwsWorkspaces#token}
	Token *MwsWorkspacesToken `field:"optional" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#workspace_id MwsWorkspaces#workspace_id}.
	WorkspaceId *float64 `field:"optional" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#workspace_status MwsWorkspaces#workspace_status}.
	WorkspaceStatus *string `field:"optional" json:"workspaceStatus" yaml:"workspaceStatus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#workspace_status_message MwsWorkspaces#workspace_status_message}.
	WorkspaceStatusMessage *string `field:"optional" json:"workspaceStatusMessage" yaml:"workspaceStatusMessage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_workspaces#workspace_url MwsWorkspaces#workspace_url}.
	WorkspaceUrl *string `field:"optional" json:"workspaceUrl" yaml:"workspaceUrl"`
}

