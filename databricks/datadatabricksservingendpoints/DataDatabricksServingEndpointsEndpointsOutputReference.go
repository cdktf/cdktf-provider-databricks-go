// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksservingendpoints/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksServingEndpointsEndpointsOutputReference interface {
	cdktf.ComplexObject
	AiGateway() DataDatabricksServingEndpointsEndpointsAiGatewayList
	AiGatewayInput() interface{}
	BudgetPolicyId() *string
	SetBudgetPolicyId(val *string)
	BudgetPolicyIdInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Config() DataDatabricksServingEndpointsEndpointsConfigList
	ConfigInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CreationTimestamp() *float64
	SetCreationTimestamp(val *float64)
	CreationTimestampInput() *float64
	Creator() *string
	SetCreator(val *string)
	CreatorInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LastUpdatedTimestamp() *float64
	SetLastUpdatedTimestamp(val *float64)
	LastUpdatedTimestampInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	State() DataDatabricksServingEndpointsEndpointsStateList
	StateInput() interface{}
	Tags() DataDatabricksServingEndpointsEndpointsTagsList
	TagsInput() interface{}
	Task() *string
	SetTask(val *string)
	TaskInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsagePolicyId() *string
	SetUsagePolicyId(val *string)
	UsagePolicyIdInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAiGateway(value interface{})
	PutConfig(value interface{})
	PutState(value interface{})
	PutTags(value interface{})
	ResetAiGateway()
	ResetBudgetPolicyId()
	ResetConfig()
	ResetCreationTimestamp()
	ResetCreator()
	ResetDescription()
	ResetId()
	ResetLastUpdatedTimestamp()
	ResetName()
	ResetState()
	ResetTags()
	ResetTask()
	ResetUsagePolicyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksServingEndpointsEndpointsOutputReference
type jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) AiGateway() DataDatabricksServingEndpointsEndpointsAiGatewayList {
	var returns DataDatabricksServingEndpointsEndpointsAiGatewayList
	_jsii_.Get(
		j,
		"aiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) AiGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aiGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) BudgetPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) BudgetPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) Config() DataDatabricksServingEndpointsEndpointsConfigList {
	var returns DataDatabricksServingEndpointsEndpointsConfigList
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) CreationTimestamp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) CreationTimestampInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) CreatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) LastUpdatedTimestamp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastUpdatedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) LastUpdatedTimestampInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastUpdatedTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) State() DataDatabricksServingEndpointsEndpointsStateList {
	var returns DataDatabricksServingEndpointsEndpointsStateList
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) StateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) Tags() DataDatabricksServingEndpointsEndpointsTagsList {
	var returns DataDatabricksServingEndpointsEndpointsTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) Task() *string {
	var returns *string
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) TaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) UsagePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) UsagePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyIdInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksServingEndpointsEndpointsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksServingEndpointsEndpointsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksServingEndpointsEndpointsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksServingEndpointsEndpointsOutputReference_Override(d DataDatabricksServingEndpointsEndpointsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetBudgetPolicyId(val *string) {
	if err := j.validateSetBudgetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetPolicyId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetCreationTimestamp(val *float64) {
	if err := j.validateSetCreationTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationTimestamp",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetCreator(val *string) {
	if err := j.validateSetCreatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creator",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetLastUpdatedTimestamp(val *float64) {
	if err := j.validateSetLastUpdatedTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastUpdatedTimestamp",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetTask(val *string) {
	if err := j.validateSetTaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"task",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference)SetUsagePolicyId(val *string) {
	if err := j.validateSetUsagePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usagePolicyId",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) PutAiGateway(value interface{}) {
	if err := d.validatePutAiGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAiGateway",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) PutConfig(value interface{}) {
	if err := d.validatePutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) PutState(value interface{}) {
	if err := d.validatePutStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putState",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetAiGateway() {
	_jsii_.InvokeVoid(
		d,
		"resetAiGateway",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetBudgetPolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetBudgetPolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetCreationTimestamp() {
	_jsii_.InvokeVoid(
		d,
		"resetCreationTimestamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetCreator() {
	_jsii_.InvokeVoid(
		d,
		"resetCreator",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetLastUpdatedTimestamp() {
	_jsii_.InvokeVoid(
		d,
		"resetLastUpdatedTimestamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		d,
		"resetState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetTask() {
	_jsii_.InvokeVoid(
		d,
		"resetTask",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ResetUsagePolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetUsagePolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

