// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksalertv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/alert_v2 databricks_alert_v2}.
type DataDatabricksAlertV2 interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	EffectiveRunAs() DataDatabricksAlertV2EffectiveRunAsOutputReference
	Evaluation() DataDatabricksAlertV2EvaluationOutputReference
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
	QueryText() *string
	SetQueryText(val *string)
	QueryTextInput() *string
	// Experimental.
	RawOverrides() interface{}
	RunAs() DataDatabricksAlertV2RunAsOutputReference
	RunAsInput() interface{}
	RunAsUserName() *string
	SetRunAsUserName(val *string)
	RunAsUserNameInput() *string
	Schedule() DataDatabricksAlertV2ScheduleOutputReference
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
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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
	PutEvaluation(value *DataDatabricksAlertV2Evaluation)
	PutRunAs(value *DataDatabricksAlertV2RunAs)
	PutSchedule(value *DataDatabricksAlertV2Schedule)
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
	ResetWorkspaceId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataDatabricksAlertV2
type jsiiProxy_DataDatabricksAlertV2 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksAlertV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) CustomDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) CustomDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) CustomSummary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) CustomSummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSummaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) EffectiveRunAs() DataDatabricksAlertV2EffectiveRunAsOutputReference {
	var returns DataDatabricksAlertV2EffectiveRunAsOutputReference
	_jsii_.Get(
		j,
		"effectiveRunAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) Evaluation() DataDatabricksAlertV2EvaluationOutputReference {
	var returns DataDatabricksAlertV2EvaluationOutputReference
	_jsii_.Get(
		j,
		"evaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) EvaluationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) OwnerUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) ParentPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) ParentPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) QueryText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) QueryTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) RunAs() DataDatabricksAlertV2RunAsOutputReference {
	var returns DataDatabricksAlertV2RunAsOutputReference
	_jsii_.Get(
		j,
		"runAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) RunAsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) RunAsUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) RunAsUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) Schedule() DataDatabricksAlertV2ScheduleOutputReference {
	var returns DataDatabricksAlertV2ScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertV2) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/alert_v2 databricks_alert_v2} Data Source.
func NewDataDatabricksAlertV2(scope constructs.Construct, id *string, config *DataDatabricksAlertV2Config) DataDatabricksAlertV2 {
	_init_.Initialize()

	if err := validateNewDataDatabricksAlertV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAlertV2{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAlertV2.DataDatabricksAlertV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/alert_v2 databricks_alert_v2} Data Source.
func NewDataDatabricksAlertV2_Override(d DataDatabricksAlertV2, scope constructs.Construct, id *string, config *DataDatabricksAlertV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAlertV2.DataDatabricksAlertV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetCustomDescription(val *string) {
	if err := j.validateSetCustomDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDescription",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetCustomSummary(val *string) {
	if err := j.validateSetCustomSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSummary",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetParentPath(val *string) {
	if err := j.validateSetParentPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentPath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetQueryText(val *string) {
	if err := j.validateSetQueryTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryText",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetRunAsUserName(val *string) {
	if err := j.validateSetRunAsUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUserName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetWarehouseId(val *string) {
	if err := j.validateSetWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertV2)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Generates CDKTF code for importing a DataDatabricksAlertV2 resource upon running "cdktf plan <stack-name>".
func DataDatabricksAlertV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatabricksAlertV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksAlertV2.DataDatabricksAlertV2",
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
func DataDatabricksAlertV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksAlertV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksAlertV2.DataDatabricksAlertV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksAlertV2_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksAlertV2_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksAlertV2.DataDatabricksAlertV2",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksAlertV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksAlertV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksAlertV2.DataDatabricksAlertV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksAlertV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksAlertV2.DataDatabricksAlertV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) PutEvaluation(value *DataDatabricksAlertV2Evaluation) {
	if err := d.validatePutEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEvaluation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) PutRunAs(value *DataDatabricksAlertV2RunAs) {
	if err := d.validatePutRunAsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRunAs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) PutSchedule(value *DataDatabricksAlertV2Schedule) {
	if err := d.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetCustomDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetCustomSummary() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomSummary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetEvaluation() {
	_jsii_.InvokeVoid(
		d,
		"resetEvaluation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetParentPath() {
	_jsii_.InvokeVoid(
		d,
		"resetParentPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetQueryText() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetRunAs() {
	_jsii_.InvokeVoid(
		d,
		"resetRunAs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetRunAsUserName() {
	_jsii_.InvokeVoid(
		d,
		"resetRunAsUserName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetWarehouseId() {
	_jsii_.InvokeVoid(
		d,
		"resetWarehouseId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) ResetWorkspaceId() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspaceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

