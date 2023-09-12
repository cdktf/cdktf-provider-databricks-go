// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwslogdelivery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v10/mwslogdelivery/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_log_delivery databricks_mws_log_delivery}.
type MwsLogDelivery interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigId() *string
	SetConfigId(val *string)
	ConfigIdInput() *string
	ConfigName() *string
	SetConfigName(val *string)
	ConfigNameInput() *string
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
	CredentialsId() *string
	SetCredentialsId(val *string)
	CredentialsIdInput() *string
	DeliveryPathPrefix() *string
	SetDeliveryPathPrefix(val *string)
	DeliveryPathPrefixInput() *string
	DeliveryStartTime() *string
	SetDeliveryStartTime(val *string)
	DeliveryStartTimeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogType() *string
	SetLogType(val *string)
	LogTypeInput() *string
	// The tree node.
	Node() constructs.Node
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
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
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	StorageConfigurationId() *string
	SetStorageConfigurationId(val *string)
	StorageConfigurationIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WorkspaceIdsFilter() *[]*float64
	SetWorkspaceIdsFilter(val *[]*float64)
	WorkspaceIdsFilterInput() *[]*float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetConfigId()
	ResetConfigName()
	ResetDeliveryPathPrefix()
	ResetDeliveryStartTime()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStatus()
	ResetWorkspaceIdsFilter()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MwsLogDelivery
type jsiiProxy_MwsLogDelivery struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MwsLogDelivery) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) ConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) ConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) ConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) ConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) CredentialsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) CredentialsIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) DeliveryPathPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) DeliveryPathPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPathPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) DeliveryStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) DeliveryStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) LogTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) StorageConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) StorageConfigurationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageConfigurationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) WorkspaceIdsFilter() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"workspaceIdsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsLogDelivery) WorkspaceIdsFilterInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"workspaceIdsFilterInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_log_delivery databricks_mws_log_delivery} Resource.
func NewMwsLogDelivery(scope constructs.Construct, id *string, config *MwsLogDeliveryConfig) MwsLogDelivery {
	_init_.Initialize()

	if err := validateNewMwsLogDeliveryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwsLogDelivery{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsLogDelivery.MwsLogDelivery",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_log_delivery databricks_mws_log_delivery} Resource.
func NewMwsLogDelivery_Override(m MwsLogDelivery, scope constructs.Construct, id *string, config *MwsLogDeliveryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsLogDelivery.MwsLogDelivery",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetConfigId(val *string) {
	if err := j.validateSetConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configId",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetConfigName(val *string) {
	if err := j.validateSetConfigNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configName",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetCredentialsId(val *string) {
	if err := j.validateSetCredentialsIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsId",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetDeliveryPathPrefix(val *string) {
	if err := j.validateSetDeliveryPathPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryPathPrefix",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetDeliveryStartTime(val *string) {
	if err := j.validateSetDeliveryStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryStartTime",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetLogType(val *string) {
	if err := j.validateSetLogTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logType",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetStorageConfigurationId(val *string) {
	if err := j.validateSetStorageConfigurationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageConfigurationId",
		val,
	)
}

func (j *jsiiProxy_MwsLogDelivery)SetWorkspaceIdsFilter(val *[]*float64) {
	if err := j.validateSetWorkspaceIdsFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceIdsFilter",
		val,
	)
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
func MwsLogDelivery_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwsLogDelivery_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.mwsLogDelivery.MwsLogDelivery",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MwsLogDelivery_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwsLogDelivery_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.mwsLogDelivery.MwsLogDelivery",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MwsLogDelivery_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwsLogDelivery_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.mwsLogDelivery.MwsLogDelivery",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MwsLogDelivery_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.mwsLogDelivery.MwsLogDelivery",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MwsLogDelivery) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MwsLogDelivery) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MwsLogDelivery) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsLogDelivery) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MwsLogDelivery) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MwsLogDelivery) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MwsLogDelivery) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MwsLogDelivery) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MwsLogDelivery) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MwsLogDelivery) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MwsLogDelivery) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsLogDelivery) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MwsLogDelivery) ResetConfigId() {
	_jsii_.InvokeVoid(
		m,
		"resetConfigId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsLogDelivery) ResetConfigName() {
	_jsii_.InvokeVoid(
		m,
		"resetConfigName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsLogDelivery) ResetDeliveryPathPrefix() {
	_jsii_.InvokeVoid(
		m,
		"resetDeliveryPathPrefix",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsLogDelivery) ResetDeliveryStartTime() {
	_jsii_.InvokeVoid(
		m,
		"resetDeliveryStartTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsLogDelivery) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsLogDelivery) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsLogDelivery) ResetStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsLogDelivery) ResetWorkspaceIdsFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkspaceIdsFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsLogDelivery) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsLogDelivery) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsLogDelivery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsLogDelivery) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

