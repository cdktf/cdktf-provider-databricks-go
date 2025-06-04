// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksexternallocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksexternallocation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference interface {
	cdktf.ComplexObject
	Algorithm() *string
	SetAlgorithm(val *string)
	AlgorithmInput() *string
	AwsKmsKeyArn() *string
	SetAwsKmsKeyArn(val *string)
	AwsKmsKeyArnInput() *string
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
	InternalValue() *DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetails
	SetInternalValue(val *DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetails)
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
	ResetAlgorithm()
	ResetAwsKmsKeyArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference
type jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) AwsKmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsKmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) AwsKmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsKmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) InternalValue() *DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetails {
	var returns *DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksExternalLocation.DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference_Override(d DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksExternalLocation.DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference)SetAwsKmsKeyArn(val *string) {
	if err := j.validateSetAwsKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsKmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference)SetInternalValue(val *DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) ResetAlgorithm() {
	_jsii_.InvokeVoid(
		d,
		"resetAlgorithm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) ResetAwsKmsKeyArn() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsKmsKeyArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEncryptionDetailsSseEncryptionDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

