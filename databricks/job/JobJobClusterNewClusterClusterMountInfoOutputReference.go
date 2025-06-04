// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobJobClusterNewClusterClusterMountInfoOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalMountDirPath() *string
	SetLocalMountDirPath(val *string)
	LocalMountDirPathInput() *string
	NetworkFilesystemInfo() JobJobClusterNewClusterClusterMountInfoNetworkFilesystemInfoOutputReference
	NetworkFilesystemInfoInput() *JobJobClusterNewClusterClusterMountInfoNetworkFilesystemInfo
	RemoteMountDirPath() *string
	SetRemoteMountDirPath(val *string)
	RemoteMountDirPathInput() *string
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
	PutNetworkFilesystemInfo(value *JobJobClusterNewClusterClusterMountInfoNetworkFilesystemInfo)
	ResetRemoteMountDirPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobJobClusterNewClusterClusterMountInfoOutputReference
type jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) LocalMountDirPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localMountDirPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) LocalMountDirPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localMountDirPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) NetworkFilesystemInfo() JobJobClusterNewClusterClusterMountInfoNetworkFilesystemInfoOutputReference {
	var returns JobJobClusterNewClusterClusterMountInfoNetworkFilesystemInfoOutputReference
	_jsii_.Get(
		j,
		"networkFilesystemInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) NetworkFilesystemInfoInput() *JobJobClusterNewClusterClusterMountInfoNetworkFilesystemInfo {
	var returns *JobJobClusterNewClusterClusterMountInfoNetworkFilesystemInfo
	_jsii_.Get(
		j,
		"networkFilesystemInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) RemoteMountDirPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteMountDirPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) RemoteMountDirPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteMountDirPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobJobClusterNewClusterClusterMountInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobJobClusterNewClusterClusterMountInfoOutputReference {
	_init_.Initialize()

	if err := validateNewJobJobClusterNewClusterClusterMountInfoOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobJobClusterNewClusterClusterMountInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobJobClusterNewClusterClusterMountInfoOutputReference_Override(j JobJobClusterNewClusterClusterMountInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobJobClusterNewClusterClusterMountInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference)SetLocalMountDirPath(val *string) {
	if err := j.validateSetLocalMountDirPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localMountDirPath",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference)SetRemoteMountDirPath(val *string) {
	if err := j.validateSetRemoteMountDirPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteMountDirPath",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) PutNetworkFilesystemInfo(value *JobJobClusterNewClusterClusterMountInfoNetworkFilesystemInfo) {
	if err := j.validatePutNetworkFilesystemInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNetworkFilesystemInfo",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) ResetRemoteMountDirPath() {
	_jsii_.InvokeVoid(
		j,
		"resetRemoteMountDirPath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobJobClusterNewClusterClusterMountInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

