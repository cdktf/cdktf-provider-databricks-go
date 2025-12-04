// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTriggerModelOutputReference interface {
	cdktf.ComplexObject
	Aliases() *[]*string
	SetAliases(val *[]*string)
	AliasesInput() *[]*string
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
	Condition() *string
	SetCondition(val *string)
	ConditionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *JobTriggerModel
	SetInternalValue(val *JobTriggerModel)
	MinTimeBetweenTriggersSeconds() *float64
	SetMinTimeBetweenTriggersSeconds(val *float64)
	MinTimeBetweenTriggersSecondsInput() *float64
	SecurableName() *string
	SetSecurableName(val *string)
	SecurableNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WaitAfterLastChangeSeconds() *float64
	SetWaitAfterLastChangeSeconds(val *float64)
	WaitAfterLastChangeSecondsInput() *float64
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
	ResetAliases()
	ResetMinTimeBetweenTriggersSeconds()
	ResetSecurableName()
	ResetWaitAfterLastChangeSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTriggerModelOutputReference
type jsiiProxy_JobTriggerModelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTriggerModelOutputReference) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) AliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) ConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) InternalValue() *JobTriggerModel {
	var returns *JobTriggerModel
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) MinTimeBetweenTriggersSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTimeBetweenTriggersSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) MinTimeBetweenTriggersSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTimeBetweenTriggersSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) SecurableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) SecurableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) WaitAfterLastChangeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitAfterLastChangeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) WaitAfterLastChangeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitAfterLastChangeSecondsInput",
		&returns,
	)
	return returns
}


func NewJobTriggerModelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTriggerModelOutputReference {
	_init_.Initialize()

	if err := validateNewJobTriggerModelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTriggerModelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTriggerModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTriggerModelOutputReference_Override(j JobTriggerModelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTriggerModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference)SetAliases(val *[]*string) {
	if err := j.validateSetAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference)SetCondition(val *string) {
	if err := j.validateSetConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference)SetInternalValue(val *JobTriggerModel) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference)SetMinTimeBetweenTriggersSeconds(val *float64) {
	if err := j.validateSetMinTimeBetweenTriggersSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTimeBetweenTriggersSeconds",
		val,
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference)SetSecurableName(val *string) {
	if err := j.validateSetSecurableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securableName",
		val,
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference)SetWaitAfterLastChangeSeconds(val *float64) {
	if err := j.validateSetWaitAfterLastChangeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitAfterLastChangeSeconds",
		val,
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTriggerModelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTriggerModelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTriggerModelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTriggerModelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTriggerModelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTriggerModelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTriggerModelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTriggerModelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTriggerModelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) ResetAliases() {
	_jsii_.InvokeVoid(
		j,
		"resetAliases",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference) ResetMinTimeBetweenTriggersSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetMinTimeBetweenTriggersSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference) ResetSecurableName() {
	_jsii_.InvokeVoid(
		j,
		"resetSecurableName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference) ResetWaitAfterLastChangeSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetWaitAfterLastChangeSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggerModelOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggerModelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

