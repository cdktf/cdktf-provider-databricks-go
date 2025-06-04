// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskForEachTaskTaskWebhookNotificationsOutputReference interface {
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
	InternalValue() *JobTaskForEachTaskTaskWebhookNotifications
	SetInternalValue(val *JobTaskForEachTaskTaskWebhookNotifications)
	OnDurationWarningThresholdExceeded() JobTaskForEachTaskTaskWebhookNotificationsOnDurationWarningThresholdExceededList
	OnDurationWarningThresholdExceededInput() interface{}
	OnFailure() JobTaskForEachTaskTaskWebhookNotificationsOnFailureList
	OnFailureInput() interface{}
	OnStart() JobTaskForEachTaskTaskWebhookNotificationsOnStartList
	OnStartInput() interface{}
	OnStreamingBacklogExceeded() JobTaskForEachTaskTaskWebhookNotificationsOnStreamingBacklogExceededList
	OnStreamingBacklogExceededInput() interface{}
	OnSuccess() JobTaskForEachTaskTaskWebhookNotificationsOnSuccessList
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
	PutOnStreamingBacklogExceeded(value interface{})
	PutOnSuccess(value interface{})
	ResetOnDurationWarningThresholdExceeded()
	ResetOnFailure()
	ResetOnStart()
	ResetOnStreamingBacklogExceeded()
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

// The jsii proxy struct for JobTaskForEachTaskTaskWebhookNotificationsOutputReference
type jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) InternalValue() *JobTaskForEachTaskTaskWebhookNotifications {
	var returns *JobTaskForEachTaskTaskWebhookNotifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) OnDurationWarningThresholdExceeded() JobTaskForEachTaskTaskWebhookNotificationsOnDurationWarningThresholdExceededList {
	var returns JobTaskForEachTaskTaskWebhookNotificationsOnDurationWarningThresholdExceededList
	_jsii_.Get(
		j,
		"onDurationWarningThresholdExceeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) OnDurationWarningThresholdExceededInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDurationWarningThresholdExceededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) OnFailure() JobTaskForEachTaskTaskWebhookNotificationsOnFailureList {
	var returns JobTaskForEachTaskTaskWebhookNotificationsOnFailureList
	_jsii_.Get(
		j,
		"onFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) OnFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) OnStart() JobTaskForEachTaskTaskWebhookNotificationsOnStartList {
	var returns JobTaskForEachTaskTaskWebhookNotificationsOnStartList
	_jsii_.Get(
		j,
		"onStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) OnStartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) OnStreamingBacklogExceeded() JobTaskForEachTaskTaskWebhookNotificationsOnStreamingBacklogExceededList {
	var returns JobTaskForEachTaskTaskWebhookNotificationsOnStreamingBacklogExceededList
	_jsii_.Get(
		j,
		"onStreamingBacklogExceeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) OnStreamingBacklogExceededInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onStreamingBacklogExceededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) OnSuccess() JobTaskForEachTaskTaskWebhookNotificationsOnSuccessList {
	var returns JobTaskForEachTaskTaskWebhookNotificationsOnSuccessList
	_jsii_.Get(
		j,
		"onSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) OnSuccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskWebhookNotificationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskWebhookNotificationsOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskWebhookNotificationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskWebhookNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskWebhookNotificationsOutputReference_Override(j JobTaskForEachTaskTaskWebhookNotificationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskWebhookNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskWebhookNotifications) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) PutOnDurationWarningThresholdExceeded(value interface{}) {
	if err := j.validatePutOnDurationWarningThresholdExceededParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnDurationWarningThresholdExceeded",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) PutOnFailure(value interface{}) {
	if err := j.validatePutOnFailureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnFailure",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) PutOnStart(value interface{}) {
	if err := j.validatePutOnStartParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnStart",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) PutOnStreamingBacklogExceeded(value interface{}) {
	if err := j.validatePutOnStreamingBacklogExceededParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnStreamingBacklogExceeded",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) PutOnSuccess(value interface{}) {
	if err := j.validatePutOnSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnSuccess",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) ResetOnDurationWarningThresholdExceeded() {
	_jsii_.InvokeVoid(
		j,
		"resetOnDurationWarningThresholdExceeded",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) ResetOnFailure() {
	_jsii_.InvokeVoid(
		j,
		"resetOnFailure",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) ResetOnStart() {
	_jsii_.InvokeVoid(
		j,
		"resetOnStart",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) ResetOnStreamingBacklogExceeded() {
	_jsii_.InvokeVoid(
		j,
		"resetOnStreamingBacklogExceeded",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) ResetOnSuccess() {
	_jsii_.InvokeVoid(
		j,
		"resetOnSuccess",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskWebhookNotificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

