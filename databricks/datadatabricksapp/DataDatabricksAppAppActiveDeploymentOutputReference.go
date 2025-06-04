// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksAppAppActiveDeploymentOutputReference interface {
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
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Creator() *string
	DeploymentArtifacts() DataDatabricksAppAppActiveDeploymentDeploymentArtifactsOutputReference
	DeploymentId() *string
	SetDeploymentId(val *string)
	DeploymentIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksAppAppActiveDeployment
	SetInternalValue(val *DataDatabricksAppAppActiveDeployment)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	SourceCodePath() *string
	SetSourceCodePath(val *string)
	SourceCodePathInput() *string
	Status() DataDatabricksAppAppActiveDeploymentStatusOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdateTime() *string
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
	ResetDeploymentId()
	ResetMode()
	ResetSourceCodePath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAppAppActiveDeploymentOutputReference
type jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) DeploymentArtifacts() DataDatabricksAppAppActiveDeploymentDeploymentArtifactsOutputReference {
	var returns DataDatabricksAppAppActiveDeploymentDeploymentArtifactsOutputReference
	_jsii_.Get(
		j,
		"deploymentArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) DeploymentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) InternalValue() *DataDatabricksAppAppActiveDeployment {
	var returns *DataDatabricksAppAppActiveDeployment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) SourceCodePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) SourceCodePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) Status() DataDatabricksAppAppActiveDeploymentStatusOutputReference {
	var returns DataDatabricksAppAppActiveDeploymentStatusOutputReference
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewDataDatabricksAppAppActiveDeploymentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksAppAppActiveDeploymentOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAppAppActiveDeploymentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksApp.DataDatabricksAppAppActiveDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAppAppActiveDeploymentOutputReference_Override(d DataDatabricksAppAppActiveDeploymentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksApp.DataDatabricksAppAppActiveDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference)SetDeploymentId(val *string) {
	if err := j.validateSetDeploymentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference)SetInternalValue(val *DataDatabricksAppAppActiveDeployment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference)SetSourceCodePath(val *string) {
	if err := j.validateSetSourceCodePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCodePath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) ResetDeploymentId() {
	_jsii_.InvokeVoid(
		d,
		"resetDeploymentId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		d,
		"resetMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) ResetSourceCodePath() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceCodePath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppActiveDeploymentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

