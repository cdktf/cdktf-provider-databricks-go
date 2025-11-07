// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscatalog/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference interface {
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
	InternalValue() *DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlag
	SetInternalValue(val *DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlag)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetInheritedFromName()
	ResetInheritedFromType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference
type jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) InheritedFromName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) InheritedFromNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) InheritedFromType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) InheritedFromTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) InternalValue() *DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlag {
	var returns *DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlag
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCatalog.DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference_Override(d DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCatalog.DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference)SetInheritedFromName(val *string) {
	if err := j.validateSetInheritedFromNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inheritedFromName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference)SetInheritedFromType(val *string) {
	if err := j.validateSetInheritedFromTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inheritedFromType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference)SetInternalValue(val *DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlag) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) ResetInheritedFromName() {
	_jsii_.InvokeVoid(
		d,
		"resetInheritedFromName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) ResetInheritedFromType() {
	_jsii_.InvokeVoid(
		d,
		"resetInheritedFromType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

