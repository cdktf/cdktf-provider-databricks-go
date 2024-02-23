// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs databricks}.
type DatabricksProvider interface {
	cdktf.TerraformProvider
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	AzureClientId() *string
	SetAzureClientId(val *string)
	AzureClientIdInput() *string
	AzureClientSecret() *string
	SetAzureClientSecret(val *string)
	AzureClientSecretInput() *string
	AzureEnvironment() *string
	SetAzureEnvironment(val *string)
	AzureEnvironmentInput() *string
	AzureLoginAppId() *string
	SetAzureLoginAppId(val *string)
	AzureLoginAppIdInput() *string
	AzureTenantId() *string
	SetAzureTenantId(val *string)
	AzureTenantIdInput() *string
	AzureUseMsi() interface{}
	SetAzureUseMsi(val interface{})
	AzureUseMsiInput() interface{}
	AzureWorkspaceResourceId() *string
	SetAzureWorkspaceResourceId(val *string)
	AzureWorkspaceResourceIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ConfigFile() *string
	SetConfigFile(val *string)
	ConfigFileInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	DatabricksCliPath() *string
	SetDatabricksCliPath(val *string)
	DatabricksCliPathInput() *string
	DebugHeaders() interface{}
	SetDebugHeaders(val interface{})
	DebugHeadersInput() interface{}
	DebugTruncateBytes() *float64
	SetDebugTruncateBytes(val *float64)
	DebugTruncateBytesInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GoogleCredentials() *string
	SetGoogleCredentials(val *string)
	GoogleCredentialsInput() *string
	GoogleServiceAccount() *string
	SetGoogleServiceAccount(val *string)
	GoogleServiceAccountInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	HttpTimeoutSeconds() *float64
	SetHttpTimeoutSeconds(val *float64)
	HttpTimeoutSecondsInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	MetadataServiceUrl() *string
	SetMetadataServiceUrl(val *string)
	MetadataServiceUrlInput() *string
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	RateLimit() *float64
	SetRateLimit(val *float64)
	RateLimitInput() *float64
	// Experimental.
	RawOverrides() interface{}
	RetryTimeoutSeconds() *float64
	SetRetryTimeoutSeconds(val *float64)
	RetryTimeoutSecondsInput() *float64
	SkipVerify() interface{}
	SetSkipVerify(val interface{})
	SkipVerifyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	WarehouseId() *string
	SetWarehouseId(val *string)
	WarehouseIdInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccountId()
	ResetAlias()
	ResetAuthType()
	ResetAzureClientId()
	ResetAzureClientSecret()
	ResetAzureEnvironment()
	ResetAzureLoginAppId()
	ResetAzureTenantId()
	ResetAzureUseMsi()
	ResetAzureWorkspaceResourceId()
	ResetClientId()
	ResetClientSecret()
	ResetClusterId()
	ResetConfigFile()
	ResetDatabricksCliPath()
	ResetDebugHeaders()
	ResetDebugTruncateBytes()
	ResetGoogleCredentials()
	ResetGoogleServiceAccount()
	ResetHost()
	ResetHttpTimeoutSeconds()
	ResetMetadataServiceUrl()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetProfile()
	ResetRateLimit()
	ResetRetryTimeoutSeconds()
	ResetSkipVerify()
	ResetToken()
	ResetUsername()
	ResetWarehouseId()
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

// The jsii proxy struct for DatabricksProvider
type jsiiProxy_DatabricksProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_DatabricksProvider) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureLoginAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureLoginAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureLoginAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureLoginAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureUseMsi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureUseMsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureUseMsiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureUseMsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureWorkspaceResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureWorkspaceResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) AzureWorkspaceResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureWorkspaceResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) ConfigFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) ConfigFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) DatabricksCliPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksCliPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) DatabricksCliPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksCliPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) DebugHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) DebugHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) DebugTruncateBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"debugTruncateBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) DebugTruncateBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"debugTruncateBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) GoogleCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) GoogleCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) GoogleServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) GoogleServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) HttpTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) HttpTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) MetadataServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) MetadataServiceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataServiceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) RateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) RateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) RetryTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) RetryTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) SkipVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) SkipVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksProvider) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs databricks} Resource.
func NewDatabricksProvider(scope constructs.Construct, id *string, config *DatabricksProviderConfig) DatabricksProvider {
	_init_.Initialize()

	if err := validateNewDatabricksProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabricksProvider{}

	_jsii_.Create(
		"@cdktf/provider-databricks.provider.DatabricksProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs databricks} Resource.
func NewDatabricksProvider_Override(d DatabricksProvider, scope constructs.Construct, id *string, config *DatabricksProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.provider.DatabricksProvider",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetAuthType(val *string) {
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetAzureClientId(val *string) {
	_jsii_.Set(
		j,
		"azureClientId",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetAzureClientSecret(val *string) {
	_jsii_.Set(
		j,
		"azureClientSecret",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetAzureEnvironment(val *string) {
	_jsii_.Set(
		j,
		"azureEnvironment",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetAzureLoginAppId(val *string) {
	_jsii_.Set(
		j,
		"azureLoginAppId",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetAzureTenantId(val *string) {
	_jsii_.Set(
		j,
		"azureTenantId",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetAzureUseMsi(val interface{}) {
	if err := j.validateSetAzureUseMsiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureUseMsi",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetAzureWorkspaceResourceId(val *string) {
	_jsii_.Set(
		j,
		"azureWorkspaceResourceId",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetConfigFile(val *string) {
	_jsii_.Set(
		j,
		"configFile",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetDatabricksCliPath(val *string) {
	_jsii_.Set(
		j,
		"databricksCliPath",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetDebugHeaders(val interface{}) {
	if err := j.validateSetDebugHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"debugHeaders",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetDebugTruncateBytes(val *float64) {
	_jsii_.Set(
		j,
		"debugTruncateBytes",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetGoogleCredentials(val *string) {
	_jsii_.Set(
		j,
		"googleCredentials",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetGoogleServiceAccount(val *string) {
	_jsii_.Set(
		j,
		"googleServiceAccount",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetHttpTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"httpTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetMetadataServiceUrl(val *string) {
	_jsii_.Set(
		j,
		"metadataServiceUrl",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetProfile(val *string) {
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetRateLimit(val *float64) {
	_jsii_.Set(
		j,
		"rateLimit",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetRetryTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"retryTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetSkipVerify(val interface{}) {
	if err := j.validateSetSkipVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipVerify",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_DatabricksProvider)SetWarehouseId(val *string) {
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

// Generates CDKTF code for importing a DatabricksProvider resource upon running "cdktf plan <stack-name>".
func DatabricksProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatabricksProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.provider.DatabricksProvider",
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
func DatabricksProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabricksProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.provider.DatabricksProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabricksProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabricksProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.provider.DatabricksProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabricksProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabricksProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.provider.DatabricksProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabricksProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.provider.DatabricksProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabricksProvider) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabricksProvider) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetAccountId() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		d,
		"resetAlias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetAuthType() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetAzureClientId() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureClientId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetAzureClientSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureClientSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetAzureEnvironment() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureEnvironment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetAzureLoginAppId() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureLoginAppId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetAzureTenantId() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureTenantId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetAzureUseMsi() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureUseMsi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetAzureWorkspaceResourceId() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureWorkspaceResourceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetClientId() {
	_jsii_.InvokeVoid(
		d,
		"resetClientId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetClientSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetConfigFile() {
	_jsii_.InvokeVoid(
		d,
		"resetConfigFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetDatabricksCliPath() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabricksCliPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetDebugHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetDebugHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetDebugTruncateBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetDebugTruncateBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetGoogleCredentials() {
	_jsii_.InvokeVoid(
		d,
		"resetGoogleCredentials",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetGoogleServiceAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetGoogleServiceAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetHost() {
	_jsii_.InvokeVoid(
		d,
		"resetHost",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetHttpTimeoutSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpTimeoutSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetMetadataServiceUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetMetadataServiceUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetRateLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetRateLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetRetryTimeoutSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetRetryTimeoutSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetSkipVerify() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipVerify",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetToken() {
	_jsii_.InvokeVoid(
		d,
		"resetToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) ResetWarehouseId() {
	_jsii_.InvokeVoid(
		d,
		"resetWarehouseId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

