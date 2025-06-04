// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference interface {
	cdktf.ComplexObject
	CleanRoomName() *string
	SetCleanRoomName(val *string)
	CleanRoomNameInput() *string
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
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskForEachTaskTaskCleanRoomsNotebookTask
	SetInternalValue(val *JobTaskForEachTaskTaskCleanRoomsNotebookTask)
	NotebookBaseParameters() *map[string]*string
	SetNotebookBaseParameters(val *map[string]*string)
	NotebookBaseParametersInput() *map[string]*string
	NotebookName() *string
	SetNotebookName(val *string)
	NotebookNameInput() *string
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
	ResetEtag()
	ResetNotebookBaseParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference
type jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) CleanRoomName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) CleanRoomNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanRoomNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) InternalValue() *JobTaskForEachTaskTaskCleanRoomsNotebookTask {
	var returns *JobTaskForEachTaskTaskCleanRoomsNotebookTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) NotebookBaseParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"notebookBaseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) NotebookBaseParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"notebookBaseParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) NotebookName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) NotebookNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference_Override(j JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference)SetCleanRoomName(val *string) {
	if err := j.validateSetCleanRoomNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanRoomName",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskCleanRoomsNotebookTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference)SetNotebookBaseParameters(val *map[string]*string) {
	if err := j.validateSetNotebookBaseParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebookBaseParameters",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference)SetNotebookName(val *string) {
	if err := j.validateSetNotebookNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebookName",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) ResetEtag() {
	_jsii_.InvokeVoid(
		j,
		"resetEtag",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) ResetNotebookBaseParameters() {
	_jsii_.InvokeVoid(
		j,
		"resetNotebookBaseParameters",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskCleanRoomsNotebookTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

