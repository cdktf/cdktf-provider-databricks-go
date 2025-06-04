// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskGenAiComputeTaskComputeOutputReference interface {
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
	GpuNodePoolId() *string
	SetGpuNodePoolId(val *string)
	GpuNodePoolIdInput() *string
	GpuType() *string
	SetGpuType(val *string)
	GpuTypeInput() *string
	InternalValue() *JobTaskGenAiComputeTaskCompute
	SetInternalValue(val *JobTaskGenAiComputeTaskCompute)
	NumGpus() *float64
	SetNumGpus(val *float64)
	NumGpusInput() *float64
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
	ResetGpuNodePoolId()
	ResetGpuType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskGenAiComputeTaskComputeOutputReference
type jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GpuNodePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuNodePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GpuNodePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuNodePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GpuType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GpuTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) InternalValue() *JobTaskGenAiComputeTaskCompute {
	var returns *JobTaskGenAiComputeTaskCompute
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) NumGpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numGpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) NumGpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numGpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskGenAiComputeTaskComputeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskGenAiComputeTaskComputeOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskGenAiComputeTaskComputeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskGenAiComputeTaskComputeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskGenAiComputeTaskComputeOutputReference_Override(j JobTaskGenAiComputeTaskComputeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskGenAiComputeTaskComputeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference)SetGpuNodePoolId(val *string) {
	if err := j.validateSetGpuNodePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuNodePoolId",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference)SetGpuType(val *string) {
	if err := j.validateSetGpuTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuType",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference)SetInternalValue(val *JobTaskGenAiComputeTaskCompute) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference)SetNumGpus(val *float64) {
	if err := j.validateSetNumGpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numGpus",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) ResetGpuNodePoolId() {
	_jsii_.InvokeVoid(
		j,
		"resetGpuNodePoolId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) ResetGpuType() {
	_jsii_.InvokeVoid(
		j,
		"resetGpuType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskGenAiComputeTaskComputeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

