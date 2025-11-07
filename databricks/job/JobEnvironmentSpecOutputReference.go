// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobEnvironmentSpecOutputReference interface {
	cdktf.ComplexObject
	Client() *string
	SetClient(val *string)
	ClientInput() *string
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
	Dependencies() *[]*string
	SetDependencies(val *[]*string)
	DependenciesInput() *[]*string
	EnvironmentVersion() *string
	SetEnvironmentVersion(val *string)
	EnvironmentVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *JobEnvironmentSpec
	SetInternalValue(val *JobEnvironmentSpec)
	JavaDependencies() *[]*string
	SetJavaDependencies(val *[]*string)
	JavaDependenciesInput() *[]*string
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
	ResetClient()
	ResetDependencies()
	ResetEnvironmentVersion()
	ResetJavaDependencies()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobEnvironmentSpecOutputReference
type jsiiProxy_JobEnvironmentSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) Client() *string {
	var returns *string
	_jsii_.Get(
		j,
		"client",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) ClientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) Dependencies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) DependenciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependenciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) EnvironmentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) EnvironmentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) InternalValue() *JobEnvironmentSpec {
	var returns *JobEnvironmentSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) JavaDependencies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"javaDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) JavaDependenciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"javaDependenciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobEnvironmentSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobEnvironmentSpecOutputReference {
	_init_.Initialize()

	if err := validateNewJobEnvironmentSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobEnvironmentSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobEnvironmentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobEnvironmentSpecOutputReference_Override(j JobEnvironmentSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobEnvironmentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference)SetClient(val *string) {
	if err := j.validateSetClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"client",
		val,
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference)SetDependencies(val *[]*string) {
	if err := j.validateSetDependenciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependencies",
		val,
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference)SetEnvironmentVersion(val *string) {
	if err := j.validateSetEnvironmentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVersion",
		val,
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference)SetInternalValue(val *JobEnvironmentSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference)SetJavaDependencies(val *[]*string) {
	if err := j.validateSetJavaDependenciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaDependencies",
		val,
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) ResetClient() {
	_jsii_.InvokeVoid(
		j,
		"resetClient",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) ResetDependencies() {
	_jsii_.InvokeVoid(
		j,
		"resetDependencies",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) ResetEnvironmentVersion() {
	_jsii_.InvokeVoid(
		j,
		"resetEnvironmentVersion",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) ResetJavaDependencies() {
	_jsii_.InvokeVoid(
		j,
		"resetJavaDependencies",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobEnvironmentSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

