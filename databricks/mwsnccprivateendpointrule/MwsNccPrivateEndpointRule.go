// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnccprivateendpointrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/mwsnccprivateendpointrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.43.0/docs/resources/mws_ncc_private_endpoint_rule databricks_mws_ncc_private_endpoint_rule}.
type MwsNccPrivateEndpointRule interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionState() *string
	SetConnectionState(val *string)
	ConnectionStateInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *float64
	SetCreationTime(val *float64)
	CreationTimeInput() *float64
	Deactivated() interface{}
	SetDeactivated(val interface{})
	DeactivatedAt() *float64
	SetDeactivatedAt(val *float64)
	DeactivatedAtInput() *float64
	DeactivatedInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndpointName() *string
	SetEndpointName(val *string)
	EndpointNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupId() *string
	SetGroupId(val *string)
	GroupIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NetworkConnectivityConfigId() *string
	SetNetworkConnectivityConfigId(val *string)
	NetworkConnectivityConfigIdInput() *string
	// The tree node.
	Node() constructs.Node
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
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	RuleId() *string
	SetRuleId(val *string)
	RuleIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedTime() *float64
	SetUpdatedTime(val *float64)
	UpdatedTimeInput() *float64
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
	ResetConnectionState()
	ResetCreationTime()
	ResetDeactivated()
	ResetDeactivatedAt()
	ResetEndpointName()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRuleId()
	ResetUpdatedTime()
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

// The jsii proxy struct for MwsNccPrivateEndpointRule
type jsiiProxy_MwsNccPrivateEndpointRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) ConnectionState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) ConnectionStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) CreationTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) Deactivated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) DeactivatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deactivatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) DeactivatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deactivatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) DeactivatedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) EndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) GroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) NetworkConnectivityConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) NetworkConnectivityConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) RuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) RuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) UpdatedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule) UpdatedTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedTimeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.43.0/docs/resources/mws_ncc_private_endpoint_rule databricks_mws_ncc_private_endpoint_rule} Resource.
func NewMwsNccPrivateEndpointRule(scope constructs.Construct, id *string, config *MwsNccPrivateEndpointRuleConfig) MwsNccPrivateEndpointRule {
	_init_.Initialize()

	if err := validateNewMwsNccPrivateEndpointRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwsNccPrivateEndpointRule{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsNccPrivateEndpointRule.MwsNccPrivateEndpointRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.43.0/docs/resources/mws_ncc_private_endpoint_rule databricks_mws_ncc_private_endpoint_rule} Resource.
func NewMwsNccPrivateEndpointRule_Override(m MwsNccPrivateEndpointRule, scope constructs.Construct, id *string, config *MwsNccPrivateEndpointRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsNccPrivateEndpointRule.MwsNccPrivateEndpointRule",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetConnectionState(val *string) {
	if err := j.validateSetConnectionStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionState",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetCreationTime(val *float64) {
	if err := j.validateSetCreationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationTime",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetDeactivated(val interface{}) {
	if err := j.validateSetDeactivatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deactivated",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetDeactivatedAt(val *float64) {
	if err := j.validateSetDeactivatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deactivatedAt",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetEndpointName(val *string) {
	if err := j.validateSetEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetGroupId(val *string) {
	if err := j.validateSetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetNetworkConnectivityConfigId(val *string) {
	if err := j.validateSetNetworkConnectivityConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConnectivityConfigId",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetRuleId(val *string) {
	if err := j.validateSetRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleId",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRule)SetUpdatedTime(val *float64) {
	if err := j.validateSetUpdatedTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedTime",
		val,
	)
}

// Generates CDKTF code for importing a MwsNccPrivateEndpointRule resource upon running "cdktf plan <stack-name>".
func MwsNccPrivateEndpointRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMwsNccPrivateEndpointRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.mwsNccPrivateEndpointRule.MwsNccPrivateEndpointRule",
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
func MwsNccPrivateEndpointRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwsNccPrivateEndpointRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.mwsNccPrivateEndpointRule.MwsNccPrivateEndpointRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MwsNccPrivateEndpointRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwsNccPrivateEndpointRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.mwsNccPrivateEndpointRule.MwsNccPrivateEndpointRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MwsNccPrivateEndpointRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwsNccPrivateEndpointRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.mwsNccPrivateEndpointRule.MwsNccPrivateEndpointRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MwsNccPrivateEndpointRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.mwsNccPrivateEndpointRule.MwsNccPrivateEndpointRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MwsNccPrivateEndpointRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsNccPrivateEndpointRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MwsNccPrivateEndpointRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MwsNccPrivateEndpointRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MwsNccPrivateEndpointRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MwsNccPrivateEndpointRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MwsNccPrivateEndpointRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MwsNccPrivateEndpointRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MwsNccPrivateEndpointRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsNccPrivateEndpointRule) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ResetConnectionState() {
	_jsii_.InvokeVoid(
		m,
		"resetConnectionState",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ResetCreationTime() {
	_jsii_.InvokeVoid(
		m,
		"resetCreationTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ResetDeactivated() {
	_jsii_.InvokeVoid(
		m,
		"resetDeactivated",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ResetDeactivatedAt() {
	_jsii_.InvokeVoid(
		m,
		"resetDeactivatedAt",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ResetEndpointName() {
	_jsii_.InvokeVoid(
		m,
		"resetEndpointName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ResetRuleId() {
	_jsii_.InvokeVoid(
		m,
		"resetRuleId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ResetUpdatedTime() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdatedTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

