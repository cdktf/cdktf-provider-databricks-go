// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference interface {
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
	InternalValue() *JobTaskForEachTaskTaskGenAiComputeTaskCompute
	SetInternalValue(val *JobTaskForEachTaskTaskGenAiComputeTaskCompute)
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

// The jsii proxy struct for JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference
type jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GpuNodePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuNodePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GpuNodePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuNodePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GpuType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GpuTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) InternalValue() *JobTaskForEachTaskTaskGenAiComputeTaskCompute {
	var returns *JobTaskForEachTaskTaskGenAiComputeTaskCompute
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) NumGpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numGpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) NumGpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numGpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference_Override(j JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference)SetGpuNodePoolId(val *string) {
	if err := j.validateSetGpuNodePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuNodePoolId",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference)SetGpuType(val *string) {
	if err := j.validateSetGpuTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuType",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskGenAiComputeTaskCompute) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference)SetNumGpus(val *float64) {
	if err := j.validateSetNumGpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numGpus",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) ResetGpuNodePoolId() {
	_jsii_.InvokeVoid(
		j,
		"resetGpuNodePoolId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) ResetGpuType() {
	_jsii_.InvokeVoid(
		j,
		"resetGpuType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskGenAiComputeTaskComputeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

