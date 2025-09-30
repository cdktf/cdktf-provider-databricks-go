// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/alertv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/alert_v2 databricks_alert_v2}.
type AlertV2 interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreateTime() *string
	CustomDescription() *string
	SetCustomDescription(val *string)
	CustomDescriptionInput() *string
	CustomSummary() *string
	SetCustomSummary(val *string)
	CustomSummaryInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveRunAs() AlertV2EffectiveRunAsOutputReference
	Evaluation() AlertV2EvaluationOutputReference
	EvaluationInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecycleState() *string
	// The tree node.
	Node() constructs.Node
	OwnerUserName() *string
	ParentPath() *string
	SetParentPath(val *string)
	ParentPathInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueryText() *string
	SetQueryText(val *string)
	QueryTextInput() *string
	// Experimental.
	RawOverrides() interface{}
	RunAs() AlertV2RunAsOutputReference
	RunAsInput() interface{}
	RunAsUserName() *string
	SetRunAsUserName(val *string)
	RunAsUserNameInput() *string
	Schedule() AlertV2ScheduleOutputReference
	ScheduleInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdateTime() *string
	WarehouseId() *string
	SetWarehouseId(val *string)
	WarehouseIdInput() *string
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
	PutEvaluation(value *AlertV2Evaluation)
	PutRunAs(value *AlertV2RunAs)
	PutSchedule(value *AlertV2Schedule)
	ResetCustomDescription()
	ResetCustomSummary()
	ResetDisplayName()
	ResetEvaluation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentPath()
	ResetQueryText()
	ResetRunAs()
	ResetRunAsUserName()
	ResetSchedule()
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

// The jsii proxy struct for AlertV2
type jsiiProxy_AlertV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AlertV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) CustomDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) CustomDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) CustomSummary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) CustomSummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSummaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) EffectiveRunAs() AlertV2EffectiveRunAsOutputReference {
	var returns AlertV2EffectiveRunAsOutputReference
	_jsii_.Get(
		j,
		"effectiveRunAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) Evaluation() AlertV2EvaluationOutputReference {
	var returns AlertV2EvaluationOutputReference
	_jsii_.Get(
		j,
		"evaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) EvaluationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) OwnerUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) ParentPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) ParentPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) QueryText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) QueryTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) RunAs() AlertV2RunAsOutputReference {
	var returns AlertV2RunAsOutputReference
	_jsii_.Get(
		j,
		"runAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) RunAsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) RunAsUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) RunAsUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) Schedule() AlertV2ScheduleOutputReference {
	var returns AlertV2ScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/alert_v2 databricks_alert_v2} Resource.
func NewAlertV2(scope constructs.Construct, id *string, config *AlertV2Config) AlertV2 {
	_init_.Initialize()

	if err := validateNewAlertV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertV2{}

	_jsii_.Create(
		"@cdktf/provider-databricks.alertV2.AlertV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/alert_v2 databricks_alert_v2} Resource.
func NewAlertV2_Override(a AlertV2, scope constructs.Construct, id *string, config *AlertV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.alertV2.AlertV2",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AlertV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetCustomDescription(val *string) {
	if err := j.validateSetCustomDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDescription",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetCustomSummary(val *string) {
	if err := j.validateSetCustomSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSummary",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetParentPath(val *string) {
	if err := j.validateSetParentPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentPath",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetQueryText(val *string) {
	if err := j.validateSetQueryTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryText",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetRunAsUserName(val *string) {
	if err := j.validateSetRunAsUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUserName",
		val,
	)
}

func (j *jsiiProxy_AlertV2)SetWarehouseId(val *string) {
	if err := j.validateSetWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

// Generates CDKTF code for importing a AlertV2 resource upon running "cdktf plan <stack-name>".
func AlertV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAlertV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.alertV2.AlertV2",
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
func AlertV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlertV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.alertV2.AlertV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlertV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlertV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.alertV2.AlertV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlertV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlertV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.alertV2.AlertV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AlertV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.alertV2.AlertV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AlertV2) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AlertV2) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AlertV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AlertV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AlertV2) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AlertV2) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AlertV2) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AlertV2) PutEvaluation(value *AlertV2Evaluation) {
	if err := a.validatePutEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEvaluation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertV2) PutRunAs(value *AlertV2RunAs) {
	if err := a.validatePutRunAsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRunAs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertV2) PutSchedule(value *AlertV2Schedule) {
	if err := a.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchedule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertV2) ResetCustomDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2) ResetCustomSummary() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomSummary",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2) ResetEvaluation() {
	_jsii_.InvokeVoid(
		a,
		"resetEvaluation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2) ResetParentPath() {
	_jsii_.InvokeVoid(
		a,
		"resetParentPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2) ResetQueryText() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2) ResetRunAs() {
	_jsii_.InvokeVoid(
		a,
		"resetRunAs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2) ResetRunAsUserName() {
	_jsii_.InvokeVoid(
		a,
		"resetRunAsUserName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2) ResetSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2) ResetWarehouseId() {
	_jsii_.InvokeVoid(
		a,
		"resetWarehouseId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

