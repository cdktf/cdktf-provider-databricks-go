// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policyinfo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/policyinfo/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/policy_info databricks_policy_info}.
type PolicyInfo interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ColumnMask() PolicyInfoColumnMaskOutputReference
	ColumnMaskInput() interface{}
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	CreatedAt() *float64
	CreatedBy() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExceptPrincipals() *[]*string
	SetExceptPrincipals(val *[]*string)
	ExceptPrincipalsInput() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForSecurableType() *string
	SetForSecurableType(val *string)
	ForSecurableTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MatchColumns() PolicyInfoMatchColumnsList
	MatchColumnsInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OnSecurableFullname() *string
	SetOnSecurableFullname(val *string)
	OnSecurableFullnameInput() *string
	OnSecurableType() *string
	SetOnSecurableType(val *string)
	OnSecurableTypeInput() *string
	PolicyType() *string
	SetPolicyType(val *string)
	PolicyTypeInput() *string
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
	RowFilter() PolicyInfoRowFilterOutputReference
	RowFilterInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ToPrincipals() *[]*string
	SetToPrincipals(val *[]*string)
	ToPrincipalsInput() *[]*string
	UpdatedAt() *float64
	UpdatedBy() *string
	WhenCondition() *string
	SetWhenCondition(val *string)
	WhenConditionInput() *string
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
	PutColumnMask(value *PolicyInfoColumnMask)
	PutMatchColumns(value interface{})
	PutRowFilter(value *PolicyInfoRowFilter)
	ResetColumnMask()
	ResetComment()
	ResetExceptPrincipals()
	ResetMatchColumns()
	ResetName()
	ResetOnSecurableFullname()
	ResetOnSecurableType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRowFilter()
	ResetWhenCondition()
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

// The jsii proxy struct for PolicyInfo
type jsiiProxy_PolicyInfo struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PolicyInfo) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) ColumnMask() PolicyInfoColumnMaskOutputReference {
	var returns PolicyInfoColumnMaskOutputReference
	_jsii_.Get(
		j,
		"columnMask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) ColumnMaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnMaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) ExceptPrincipals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exceptPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) ExceptPrincipalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exceptPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) ForSecurableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forSecurableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) ForSecurableTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forSecurableTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) MatchColumns() PolicyInfoMatchColumnsList {
	var returns PolicyInfoMatchColumnsList
	_jsii_.Get(
		j,
		"matchColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) MatchColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) OnSecurableFullname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onSecurableFullname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) OnSecurableFullnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onSecurableFullnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) OnSecurableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onSecurableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) OnSecurableTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onSecurableTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) RowFilter() PolicyInfoRowFilterOutputReference {
	var returns PolicyInfoRowFilterOutputReference
	_jsii_.Get(
		j,
		"rowFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) RowFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rowFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) ToPrincipals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"toPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) ToPrincipalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"toPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) WhenCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfo) WhenConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenConditionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/policy_info databricks_policy_info} Resource.
func NewPolicyInfo(scope constructs.Construct, id *string, config *PolicyInfoConfig) PolicyInfo {
	_init_.Initialize()

	if err := validateNewPolicyInfoParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyInfo{}

	_jsii_.Create(
		"@cdktf/provider-databricks.policyInfo.PolicyInfo",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/policy_info databricks_policy_info} Resource.
func NewPolicyInfo_Override(p PolicyInfo, scope constructs.Construct, id *string, config *PolicyInfoConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.policyInfo.PolicyInfo",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PolicyInfo)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetExceptPrincipals(val *[]*string) {
	if err := j.validateSetExceptPrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exceptPrincipals",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetForSecurableType(val *string) {
	if err := j.validateSetForSecurableTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forSecurableType",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetOnSecurableFullname(val *string) {
	if err := j.validateSetOnSecurableFullnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onSecurableFullname",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetOnSecurableType(val *string) {
	if err := j.validateSetOnSecurableTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onSecurableType",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetPolicyType(val *string) {
	if err := j.validateSetPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetToPrincipals(val *[]*string) {
	if err := j.validateSetToPrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toPrincipals",
		val,
	)
}

func (j *jsiiProxy_PolicyInfo)SetWhenCondition(val *string) {
	if err := j.validateSetWhenConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whenCondition",
		val,
	)
}

// Generates CDKTF code for importing a PolicyInfo resource upon running "cdktf plan <stack-name>".
func PolicyInfo_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePolicyInfo_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.policyInfo.PolicyInfo",
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
func PolicyInfo_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyInfo_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.policyInfo.PolicyInfo",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyInfo_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyInfo_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.policyInfo.PolicyInfo",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyInfo_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyInfo_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.policyInfo.PolicyInfo",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PolicyInfo_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.policyInfo.PolicyInfo",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PolicyInfo) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PolicyInfo) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PolicyInfo) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PolicyInfo) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyInfo) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PolicyInfo) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyInfo) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PolicyInfo) PutColumnMask(value *PolicyInfoColumnMask) {
	if err := p.validatePutColumnMaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putColumnMask",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyInfo) PutMatchColumns(value interface{}) {
	if err := p.validatePutMatchColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMatchColumns",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyInfo) PutRowFilter(value *PolicyInfoRowFilter) {
	if err := p.validatePutRowFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRowFilter",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyInfo) ResetColumnMask() {
	_jsii_.InvokeVoid(
		p,
		"resetColumnMask",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfo) ResetComment() {
	_jsii_.InvokeVoid(
		p,
		"resetComment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfo) ResetExceptPrincipals() {
	_jsii_.InvokeVoid(
		p,
		"resetExceptPrincipals",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfo) ResetMatchColumns() {
	_jsii_.InvokeVoid(
		p,
		"resetMatchColumns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfo) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfo) ResetOnSecurableFullname() {
	_jsii_.InvokeVoid(
		p,
		"resetOnSecurableFullname",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfo) ResetOnSecurableType() {
	_jsii_.InvokeVoid(
		p,
		"resetOnSecurableType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfo) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfo) ResetRowFilter() {
	_jsii_.InvokeVoid(
		p,
		"resetRowFilter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfo) ResetWhenCondition() {
	_jsii_.InvokeVoid(
		p,
		"resetWhenCondition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfo) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfo) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

