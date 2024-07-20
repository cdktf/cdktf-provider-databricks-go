// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksschema

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabricksschema/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference interface {
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
	InheritedFromName() *string
	SetInheritedFromName(val *string)
	InheritedFromNameInput() *string
	InheritedFromType() *string
	SetInheritedFromType(val *string)
	InheritedFromTypeInput() *string
	InternalValue() *DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlag
	SetInternalValue(val *DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlag)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetInheritedFromName()
	ResetInheritedFromType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference
type jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) InheritedFromName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) InheritedFromNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) InheritedFromType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) InheritedFromTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) InternalValue() *DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlag {
	var returns *DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlag
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksSchema.DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference_Override(d DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksSchema.DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference)SetInheritedFromName(val *string) {
	if err := j.validateSetInheritedFromNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inheritedFromName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference)SetInheritedFromType(val *string) {
	if err := j.validateSetInheritedFromTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inheritedFromType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference)SetInternalValue(val *DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlag) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) ResetInheritedFromName() {
	_jsii_.InvokeVoid(
		d,
		"resetInheritedFromName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) ResetInheritedFromType() {
	_jsii_.InvokeVoid(
		d,
		"resetInheritedFromType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

