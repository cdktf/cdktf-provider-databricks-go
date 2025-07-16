// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsworkspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/mwsworkspaces/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/mws_workspaces databricks_mws_workspaces}.
type MwsWorkspaces interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AwsRegion() *string
	SetAwsRegion(val *string)
	AwsRegionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cloud() *string
	SetCloud(val *string)
	CloudInput() *string
	CloudResourceContainer() MwsWorkspacesCloudResourceContainerOutputReference
	CloudResourceContainerInput() *MwsWorkspacesCloudResourceContainer
	ComputeMode() *string
	SetComputeMode(val *string)
	ComputeModeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *float64
	SetCreationTime(val *float64)
	CreationTimeInput() *float64
	CredentialsId() *string
	SetCredentialsId(val *string)
	CredentialsIdInput() *string
	CustomerManagedKeyId() *string
	SetCustomerManagedKeyId(val *string)
	CustomerManagedKeyIdInput() *string
	CustomTags() *map[string]*string
	SetCustomTags(val *map[string]*string)
	CustomTagsInput() *map[string]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentName() *string
	SetDeploymentName(val *string)
	DeploymentNameInput() *string
	EffectiveComputeMode() *string
	ExternalCustomerInfo() MwsWorkspacesExternalCustomerInfoOutputReference
	ExternalCustomerInfoInput() *MwsWorkspacesExternalCustomerInfo
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcpManagedNetworkConfig() MwsWorkspacesGcpManagedNetworkConfigOutputReference
	GcpManagedNetworkConfigInput() *MwsWorkspacesGcpManagedNetworkConfig
	GcpWorkspaceSa() *string
	GkeConfig() MwsWorkspacesGkeConfigOutputReference
	GkeConfigInput() *MwsWorkspacesGkeConfig
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsNoPublicIpEnabled() interface{}
	SetIsNoPublicIpEnabled(val interface{})
	IsNoPublicIpEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagedServicesCustomerManagedKeyId() *string
	SetManagedServicesCustomerManagedKeyId(val *string)
	ManagedServicesCustomerManagedKeyIdInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
	// The tree node.
	Node() constructs.Node
	PricingTier() *string
	SetPricingTier(val *string)
	PricingTierInput() *string
	PrivateAccessSettingsId() *string
	SetPrivateAccessSettingsId(val *string)
	PrivateAccessSettingsIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	StorageConfigurationId() *string
	SetStorageConfigurationId(val *string)
	StorageConfigurationIdInput() *string
	StorageCustomerManagedKeyId() *string
	SetStorageCustomerManagedKeyId(val *string)
	StorageCustomerManagedKeyIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MwsWorkspacesTimeoutsOutputReference
	TimeoutsInput() interface{}
	Token() MwsWorkspacesTokenOutputReference
	TokenInput() *MwsWorkspacesToken
	WorkspaceId() *float64
	SetWorkspaceId(val *float64)
	WorkspaceIdInput() *float64
	WorkspaceName() *string
	SetWorkspaceName(val *string)
	WorkspaceNameInput() *string
	WorkspaceStatus() *string
	SetWorkspaceStatus(val *string)
	WorkspaceStatusInput() *string
	WorkspaceStatusMessage() *string
	SetWorkspaceStatusMessage(val *string)
	WorkspaceStatusMessageInput() *string
	WorkspaceUrl() *string
	SetWorkspaceUrl(val *string)
	WorkspaceUrlInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCloudResourceContainer(value *MwsWorkspacesCloudResourceContainer)
	PutExternalCustomerInfo(value *MwsWorkspacesExternalCustomerInfo)
	PutGcpManagedNetworkConfig(value *MwsWorkspacesGcpManagedNetworkConfig)
	PutGkeConfig(value *MwsWorkspacesGkeConfig)
	PutTimeouts(value *MwsWorkspacesTimeouts)
	PutToken(value *MwsWorkspacesToken)
	ResetAwsRegion()
	ResetCloud()
	ResetCloudResourceContainer()
	ResetComputeMode()
	ResetCreationTime()
	ResetCredentialsId()
	ResetCustomerManagedKeyId()
	ResetCustomTags()
	ResetDeploymentName()
	ResetExternalCustomerInfo()
	ResetGcpManagedNetworkConfig()
	ResetGkeConfig()
	ResetId()
	ResetIsNoPublicIpEnabled()
	ResetLocation()
	ResetManagedServicesCustomerManagedKeyId()
	ResetNetworkId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPricingTier()
	ResetPrivateAccessSettingsId()
	ResetStorageConfigurationId()
	ResetStorageCustomerManagedKeyId()
	ResetTimeouts()
	ResetToken()
	ResetWorkspaceId()
	ResetWorkspaceStatus()
	ResetWorkspaceStatusMessage()
	ResetWorkspaceUrl()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MwsWorkspaces
type jsiiProxy_MwsWorkspaces struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MwsWorkspaces) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Cloud() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CloudInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CloudResourceContainer() MwsWorkspacesCloudResourceContainerOutputReference {
	var returns MwsWorkspacesCloudResourceContainerOutputReference
	_jsii_.Get(
		j,
		"cloudResourceContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CloudResourceContainerInput() *MwsWorkspacesCloudResourceContainer {
	var returns *MwsWorkspacesCloudResourceContainer
	_jsii_.Get(
		j,
		"cloudResourceContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) ComputeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) ComputeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CreationTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CredentialsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CredentialsIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CustomerManagedKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CustomerManagedKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) DeploymentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) DeploymentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) EffectiveComputeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveComputeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) ExternalCustomerInfo() MwsWorkspacesExternalCustomerInfoOutputReference {
	var returns MwsWorkspacesExternalCustomerInfoOutputReference
	_jsii_.Get(
		j,
		"externalCustomerInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) ExternalCustomerInfoInput() *MwsWorkspacesExternalCustomerInfo {
	var returns *MwsWorkspacesExternalCustomerInfo
	_jsii_.Get(
		j,
		"externalCustomerInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) GcpManagedNetworkConfig() MwsWorkspacesGcpManagedNetworkConfigOutputReference {
	var returns MwsWorkspacesGcpManagedNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"gcpManagedNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) GcpManagedNetworkConfigInput() *MwsWorkspacesGcpManagedNetworkConfig {
	var returns *MwsWorkspacesGcpManagedNetworkConfig
	_jsii_.Get(
		j,
		"gcpManagedNetworkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) GcpWorkspaceSa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpWorkspaceSa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) GkeConfig() MwsWorkspacesGkeConfigOutputReference {
	var returns MwsWorkspacesGkeConfigOutputReference
	_jsii_.Get(
		j,
		"gkeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) GkeConfigInput() *MwsWorkspacesGkeConfig {
	var returns *MwsWorkspacesGkeConfig
	_jsii_.Get(
		j,
		"gkeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) IsNoPublicIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNoPublicIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) IsNoPublicIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNoPublicIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) ManagedServicesCustomerManagedKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedServicesCustomerManagedKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) ManagedServicesCustomerManagedKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedServicesCustomerManagedKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) PricingTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pricingTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) PricingTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pricingTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) PrivateAccessSettingsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateAccessSettingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) PrivateAccessSettingsIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateAccessSettingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) StorageConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) StorageConfigurationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageConfigurationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) StorageCustomerManagedKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCustomerManagedKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) StorageCustomerManagedKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCustomerManagedKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Timeouts() MwsWorkspacesTimeoutsOutputReference {
	var returns MwsWorkspacesTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) Token() MwsWorkspacesTokenOutputReference {
	var returns MwsWorkspacesTokenOutputReference
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) TokenInput() *MwsWorkspacesToken {
	var returns *MwsWorkspacesToken
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) WorkspaceId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) WorkspaceIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) WorkspaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) WorkspaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) WorkspaceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) WorkspaceStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) WorkspaceStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) WorkspaceStatusMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceStatusMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) WorkspaceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspaces) WorkspaceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceUrlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/mws_workspaces databricks_mws_workspaces} Resource.
func NewMwsWorkspaces(scope constructs.Construct, id *string, config *MwsWorkspacesConfig) MwsWorkspaces {
	_init_.Initialize()

	if err := validateNewMwsWorkspacesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwsWorkspaces{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsWorkspaces.MwsWorkspaces",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/mws_workspaces databricks_mws_workspaces} Resource.
func NewMwsWorkspaces_Override(m MwsWorkspaces, scope constructs.Construct, id *string, config *MwsWorkspacesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsWorkspaces.MwsWorkspaces",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetAwsRegion(val *string) {
	if err := j.validateSetAwsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetCloud(val *string) {
	if err := j.validateSetCloudParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloud",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetComputeMode(val *string) {
	if err := j.validateSetComputeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeMode",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetCreationTime(val *float64) {
	if err := j.validateSetCreationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationTime",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetCredentialsId(val *string) {
	if err := j.validateSetCredentialsIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsId",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetCustomerManagedKeyId(val *string) {
	if err := j.validateSetCustomerManagedKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerManagedKeyId",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetDeploymentName(val *string) {
	if err := j.validateSetDeploymentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentName",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetIsNoPublicIpEnabled(val interface{}) {
	if err := j.validateSetIsNoPublicIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isNoPublicIpEnabled",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetManagedServicesCustomerManagedKeyId(val *string) {
	if err := j.validateSetManagedServicesCustomerManagedKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedServicesCustomerManagedKeyId",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetNetworkId(val *string) {
	if err := j.validateSetNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetPricingTier(val *string) {
	if err := j.validateSetPricingTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pricingTier",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetPrivateAccessSettingsId(val *string) {
	if err := j.validateSetPrivateAccessSettingsIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateAccessSettingsId",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetStorageConfigurationId(val *string) {
	if err := j.validateSetStorageConfigurationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageConfigurationId",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetStorageCustomerManagedKeyId(val *string) {
	if err := j.validateSetStorageCustomerManagedKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCustomerManagedKeyId",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetWorkspaceId(val *float64) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetWorkspaceName(val *string) {
	if err := j.validateSetWorkspaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceName",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetWorkspaceStatus(val *string) {
	if err := j.validateSetWorkspaceStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceStatus",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetWorkspaceStatusMessage(val *string) {
	if err := j.validateSetWorkspaceStatusMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceStatusMessage",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspaces)SetWorkspaceUrl(val *string) {
	if err := j.validateSetWorkspaceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceUrl",
		val,
	)
}

// Generates CDKTF code for importing a MwsWorkspaces resource upon running "cdktf plan <stack-name>".
func MwsWorkspaces_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMwsWorkspaces_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.mwsWorkspaces.MwsWorkspaces",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func MwsWorkspaces_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwsWorkspaces_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.mwsWorkspaces.MwsWorkspaces",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MwsWorkspaces_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwsWorkspaces_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.mwsWorkspaces.MwsWorkspaces",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MwsWorkspaces_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwsWorkspaces_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.mwsWorkspaces.MwsWorkspaces",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MwsWorkspaces_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.mwsWorkspaces.MwsWorkspaces",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MwsWorkspaces) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MwsWorkspaces) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MwsWorkspaces) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MwsWorkspaces) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MwsWorkspaces) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MwsWorkspaces) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MwsWorkspaces) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MwsWorkspaces) PutCloudResourceContainer(value *MwsWorkspacesCloudResourceContainer) {
	if err := m.validatePutCloudResourceContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCloudResourceContainer",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwsWorkspaces) PutExternalCustomerInfo(value *MwsWorkspacesExternalCustomerInfo) {
	if err := m.validatePutExternalCustomerInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putExternalCustomerInfo",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwsWorkspaces) PutGcpManagedNetworkConfig(value *MwsWorkspacesGcpManagedNetworkConfig) {
	if err := m.validatePutGcpManagedNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGcpManagedNetworkConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwsWorkspaces) PutGkeConfig(value *MwsWorkspacesGkeConfig) {
	if err := m.validatePutGkeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGkeConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwsWorkspaces) PutTimeouts(value *MwsWorkspacesTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwsWorkspaces) PutToken(value *MwsWorkspacesToken) {
	if err := m.validatePutTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putToken",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetAwsRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetAwsRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetCloud() {
	_jsii_.InvokeVoid(
		m,
		"resetCloud",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetCloudResourceContainer() {
	_jsii_.InvokeVoid(
		m,
		"resetCloudResourceContainer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetComputeMode() {
	_jsii_.InvokeVoid(
		m,
		"resetComputeMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetCreationTime() {
	_jsii_.InvokeVoid(
		m,
		"resetCreationTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetCredentialsId() {
	_jsii_.InvokeVoid(
		m,
		"resetCredentialsId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetCustomerManagedKeyId() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomerManagedKeyId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetCustomTags() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetDeploymentName() {
	_jsii_.InvokeVoid(
		m,
		"resetDeploymentName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetExternalCustomerInfo() {
	_jsii_.InvokeVoid(
		m,
		"resetExternalCustomerInfo",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetGcpManagedNetworkConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetGcpManagedNetworkConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetGkeConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetGkeConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetIsNoPublicIpEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetIsNoPublicIpEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetLocation() {
	_jsii_.InvokeVoid(
		m,
		"resetLocation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetManagedServicesCustomerManagedKeyId() {
	_jsii_.InvokeVoid(
		m,
		"resetManagedServicesCustomerManagedKeyId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetNetworkId() {
	_jsii_.InvokeVoid(
		m,
		"resetNetworkId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetPricingTier() {
	_jsii_.InvokeVoid(
		m,
		"resetPricingTier",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetPrivateAccessSettingsId() {
	_jsii_.InvokeVoid(
		m,
		"resetPrivateAccessSettingsId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetStorageConfigurationId() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageConfigurationId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetStorageCustomerManagedKeyId() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageCustomerManagedKeyId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetToken() {
	_jsii_.InvokeVoid(
		m,
		"resetToken",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetWorkspaceId() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkspaceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetWorkspaceStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkspaceStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetWorkspaceStatusMessage() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkspaceStatusMessage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) ResetWorkspaceUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkspaceUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsWorkspaces) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspaces) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

