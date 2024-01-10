// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskWebhookNotificationsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskWebhookNotifications
	SetInternalValue(val *JobTaskWebhookNotifications)
	OnDurationWarningThresholdExceeded() JobTaskWebhookNotificationsOnDurationWarningThresholdExceededList
	OnDurationWarningThresholdExceededInput() interface{}
	OnFailure() JobTaskWebhookNotificationsOnFailureList
	OnFailureInput() interface{}
	OnStart() JobTaskWebhookNotificationsOnStartList
	OnStartInput() interface{}
	OnSuccess() JobTaskWebhookNotificationsOnSuccessList
	OnSuccessInput() interface{}
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutOnDurationWarningThresholdExceeded(value interface{})
	PutOnFailure(value interface{})
	PutOnStart(value interface{})
	PutOnSuccess(value interface{})
	ResetOnDurationWarningThresholdExceeded()
	ResetOnFailure()
	ResetOnStart()
	ResetOnSuccess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskWebhookNotificationsOutputReference
type jsiiProxy_JobTaskWebhookNotificationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) InternalValue() *JobTaskWebhookNotifications {
	var returns *JobTaskWebhookNotifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) OnDurationWarningThresholdExceeded() JobTaskWebhookNotificationsOnDurationWarningThresholdExceededList {
	var returns JobTaskWebhookNotificationsOnDurationWarningThresholdExceededList
	_jsii_.Get(
		j,
		"onDurationWarningThresholdExceeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) OnDurationWarningThresholdExceededInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDurationWarningThresholdExceededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) OnFailure() JobTaskWebhookNotificationsOnFailureList {
	var returns JobTaskWebhookNotificationsOnFailureList
	_jsii_.Get(
		j,
		"onFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) OnFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) OnStart() JobTaskWebhookNotificationsOnStartList {
	var returns JobTaskWebhookNotificationsOnStartList
	_jsii_.Get(
		j,
		"onStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) OnStartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) OnSuccess() JobTaskWebhookNotificationsOnSuccessList {
	var returns JobTaskWebhookNotificationsOnSuccessList
	_jsii_.Get(
		j,
		"onSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) OnSuccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskWebhookNotificationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskWebhookNotificationsOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskWebhookNotificationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskWebhookNotificationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskWebhookNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskWebhookNotificationsOutputReference_Override(j JobTaskWebhookNotificationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskWebhookNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference)SetInternalValue(val *JobTaskWebhookNotifications) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) PutOnDurationWarningThresholdExceeded(value interface{}) {
	if err := j.validatePutOnDurationWarningThresholdExceededParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnDurationWarningThresholdExceeded",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) PutOnFailure(value interface{}) {
	if err := j.validatePutOnFailureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnFailure",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) PutOnStart(value interface{}) {
	if err := j.validatePutOnStartParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnStart",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) PutOnSuccess(value interface{}) {
	if err := j.validatePutOnSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnSuccess",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) ResetOnDurationWarningThresholdExceeded() {
	_jsii_.InvokeVoid(
		j,
		"resetOnDurationWarningThresholdExceeded",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) ResetOnFailure() {
	_jsii_.InvokeVoid(
		j,
		"resetOnFailure",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) ResetOnStart() {
	_jsii_.InvokeVoid(
		j,
		"resetOnStart",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) ResetOnSuccess() {
	_jsii_.InvokeVoid(
		j,
		"resetOnSuccess",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskWebhookNotificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

