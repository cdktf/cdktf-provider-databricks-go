// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/budget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BudgetAlertConfigurationsOutputReference interface {
	cdktf.ComplexObject
	ActionConfigurations() BudgetAlertConfigurationsActionConfigurationsList
	ActionConfigurationsInput() interface{}
	AlertConfigurationId() *string
	SetAlertConfigurationId(val *string)
	AlertConfigurationIdInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	QuantityThreshold() *string
	SetQuantityThreshold(val *string)
	QuantityThresholdInput() *string
	QuantityType() *string
	SetQuantityType(val *string)
	QuantityTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimePeriod() *string
	SetTimePeriod(val *string)
	TimePeriodInput() *string
	TriggerType() *string
	SetTriggerType(val *string)
	TriggerTypeInput() *string
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
	PutActionConfigurations(value interface{})
	ResetActionConfigurations()
	ResetAlertConfigurationId()
	ResetQuantityThreshold()
	ResetQuantityType()
	ResetTimePeriod()
	ResetTriggerType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BudgetAlertConfigurationsOutputReference
type jsiiProxy_BudgetAlertConfigurationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) ActionConfigurations() BudgetAlertConfigurationsActionConfigurationsList {
	var returns BudgetAlertConfigurationsActionConfigurationsList
	_jsii_.Get(
		j,
		"actionConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) ActionConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) AlertConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) AlertConfigurationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertConfigurationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) QuantityThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quantityThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) QuantityThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quantityThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) QuantityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quantityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) QuantityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quantityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) TimePeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) TimePeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) TriggerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference) TriggerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerTypeInput",
		&returns,
	)
	return returns
}


func NewBudgetAlertConfigurationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BudgetAlertConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewBudgetAlertConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BudgetAlertConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.budget.BudgetAlertConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBudgetAlertConfigurationsOutputReference_Override(b BudgetAlertConfigurationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.budget.BudgetAlertConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference)SetAlertConfigurationId(val *string) {
	if err := j.validateSetAlertConfigurationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertConfigurationId",
		val,
	)
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference)SetQuantityThreshold(val *string) {
	if err := j.validateSetQuantityThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quantityThreshold",
		val,
	)
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference)SetQuantityType(val *string) {
	if err := j.validateSetQuantityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quantityType",
		val,
	)
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference)SetTimePeriod(val *string) {
	if err := j.validateSetTimePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timePeriod",
		val,
	)
}

func (j *jsiiProxy_BudgetAlertConfigurationsOutputReference)SetTriggerType(val *string) {
	if err := j.validateSetTriggerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerType",
		val,
	)
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) PutActionConfigurations(value interface{}) {
	if err := b.validatePutActionConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putActionConfigurations",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) ResetActionConfigurations() {
	_jsii_.InvokeVoid(
		b,
		"resetActionConfigurations",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) ResetAlertConfigurationId() {
	_jsii_.InvokeVoid(
		b,
		"resetAlertConfigurationId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) ResetQuantityThreshold() {
	_jsii_.InvokeVoid(
		b,
		"resetQuantityThreshold",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) ResetQuantityType() {
	_jsii_.InvokeVoid(
		b,
		"resetQuantityType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) ResetTimePeriod() {
	_jsii_.InvokeVoid(
		b,
		"resetTimePeriod",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) ResetTriggerType() {
	_jsii_.InvokeVoid(
		b,
		"resetTriggerType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetAlertConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

