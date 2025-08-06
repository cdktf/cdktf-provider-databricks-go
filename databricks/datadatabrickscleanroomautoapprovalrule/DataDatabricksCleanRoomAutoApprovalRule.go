// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomautoapprovalrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscleanroomautoapprovalrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_auto_approval_rule databricks_clean_room_auto_approval_rule}.
type DataDatabricksCleanRoomAutoApprovalRule interface {
	cdktf.TerraformDataSource
	AuthorCollaboratorAlias() *string
	SetAuthorCollaboratorAlias(val *string)
	AuthorCollaboratorAliasInput() *string
	AuthorScope() *string
	SetAuthorScope(val *string)
	AuthorScopeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CleanRoomName() *string
	SetCleanRoomName(val *string)
	CleanRoomNameInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *float64
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RuleId() *string
	RuleOwnerCollaboratorAlias() *string
	RunnerCollaboratorAlias() *string
	SetRunnerCollaboratorAlias(val *string)
	RunnerCollaboratorAliasInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAuthorCollaboratorAlias()
	ResetAuthorScope()
	ResetCleanRoomName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRunnerCollaboratorAlias()
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

// The jsii proxy struct for DataDatabricksCleanRoomAutoApprovalRule
type jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) AuthorCollaboratorAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorCollaboratorAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) AuthorCollaboratorAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorCollaboratorAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) AuthorScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) AuthorScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) CleanRoomName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) CleanRoomNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) RuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) RuleOwnerCollaboratorAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleOwnerCollaboratorAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) RunnerCollaboratorAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnerCollaboratorAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) RunnerCollaboratorAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnerCollaboratorAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_auto_approval_rule databricks_clean_room_auto_approval_rule} Data Source.
func NewDataDatabricksCleanRoomAutoApprovalRule(scope constructs.Construct, id *string, config *DataDatabricksCleanRoomAutoApprovalRuleConfig) DataDatabricksCleanRoomAutoApprovalRule {
	_init_.Initialize()

	if err := validateNewDataDatabricksCleanRoomAutoApprovalRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAutoApprovalRule.DataDatabricksCleanRoomAutoApprovalRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/clean_room_auto_approval_rule databricks_clean_room_auto_approval_rule} Data Source.
func NewDataDatabricksCleanRoomAutoApprovalRule_Override(d DataDatabricksCleanRoomAutoApprovalRule, scope constructs.Construct, id *string, config *DataDatabricksCleanRoomAutoApprovalRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAutoApprovalRule.DataDatabricksCleanRoomAutoApprovalRule",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule)SetAuthorCollaboratorAlias(val *string) {
	if err := j.validateSetAuthorCollaboratorAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorCollaboratorAlias",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule)SetAuthorScope(val *string) {
	if err := j.validateSetAuthorScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorScope",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule)SetCleanRoomName(val *string) {
	if err := j.validateSetCleanRoomNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanRoomName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule)SetRunnerCollaboratorAlias(val *string) {
	if err := j.validateSetRunnerCollaboratorAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runnerCollaboratorAlias",
		val,
	)
}

// Generates CDKTF code for importing a DataDatabricksCleanRoomAutoApprovalRule resource upon running "cdktf plan <stack-name>".
func DataDatabricksCleanRoomAutoApprovalRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatabricksCleanRoomAutoApprovalRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAutoApprovalRule.DataDatabricksCleanRoomAutoApprovalRule",
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
func DataDatabricksCleanRoomAutoApprovalRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksCleanRoomAutoApprovalRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAutoApprovalRule.DataDatabricksCleanRoomAutoApprovalRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksCleanRoomAutoApprovalRule_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksCleanRoomAutoApprovalRule_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAutoApprovalRule.DataDatabricksCleanRoomAutoApprovalRule",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksCleanRoomAutoApprovalRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksCleanRoomAutoApprovalRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAutoApprovalRule.DataDatabricksCleanRoomAutoApprovalRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksCleanRoomAutoApprovalRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomAutoApprovalRule.DataDatabricksCleanRoomAutoApprovalRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) ResetAuthorCollaboratorAlias() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthorCollaboratorAlias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) ResetAuthorScope() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthorScope",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) ResetCleanRoomName() {
	_jsii_.InvokeVoid(
		d,
		"resetCleanRoomName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) ResetRunnerCollaboratorAlias() {
	_jsii_.InvokeVoid(
		d,
		"resetRunnerCollaboratorAlias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomAutoApprovalRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

