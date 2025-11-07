// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/alertv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertV2EvaluationNotificationOutputReference interface {
	cdktf.ComplexObject
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
	EffectiveNotifyOnOk() cdktf.IResolvable
	EffectiveRetriggerSeconds() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NotifyOnOk() interface{}
	SetNotifyOnOk(val interface{})
	NotifyOnOkInput() interface{}
	RetriggerSeconds() *float64
	SetRetriggerSeconds(val *float64)
	RetriggerSecondsInput() *float64
	Subscriptions() AlertV2EvaluationNotificationSubscriptionsList
	SubscriptionsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutSubscriptions(value interface{})
	ResetNotifyOnOk()
	ResetRetriggerSeconds()
	ResetSubscriptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertV2EvaluationNotificationOutputReference
type jsiiProxy_AlertV2EvaluationNotificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) EffectiveNotifyOnOk() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"effectiveNotifyOnOk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) EffectiveRetriggerSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"effectiveRetriggerSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) NotifyOnOk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnOk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) NotifyOnOkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnOkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) RetriggerSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retriggerSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) RetriggerSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retriggerSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) Subscriptions() AlertV2EvaluationNotificationSubscriptionsList {
	var returns AlertV2EvaluationNotificationSubscriptionsList
	_jsii_.Get(
		j,
		"subscriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) SubscriptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlertV2EvaluationNotificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlertV2EvaluationNotificationOutputReference {
	_init_.Initialize()

	if err := validateNewAlertV2EvaluationNotificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertV2EvaluationNotificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.alertV2.AlertV2EvaluationNotificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlertV2EvaluationNotificationOutputReference_Override(a AlertV2EvaluationNotificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.alertV2.AlertV2EvaluationNotificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference)SetNotifyOnOk(val interface{}) {
	if err := j.validateSetNotifyOnOkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnOk",
		val,
	)
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference)SetRetriggerSeconds(val *float64) {
	if err := j.validateSetRetriggerSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retriggerSeconds",
		val,
	)
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertV2EvaluationNotificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) PutSubscriptions(value interface{}) {
	if err := a.validatePutSubscriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubscriptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) ResetNotifyOnOk() {
	_jsii_.InvokeVoid(
		a,
		"resetNotifyOnOk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) ResetRetriggerSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetRetriggerSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) ResetSubscriptions() {
	_jsii_.InvokeVoid(
		a,
		"resetSubscriptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertV2EvaluationNotificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

