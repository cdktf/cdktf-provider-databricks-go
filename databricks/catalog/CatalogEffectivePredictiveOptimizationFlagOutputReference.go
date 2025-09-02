// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package catalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/catalog/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CatalogEffectivePredictiveOptimizationFlagOutputReference interface {
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
	InternalValue() *CatalogEffectivePredictiveOptimizationFlag
	SetInternalValue(val *CatalogEffectivePredictiveOptimizationFlag)
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

// The jsii proxy struct for CatalogEffectivePredictiveOptimizationFlagOutputReference
type jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) InheritedFromName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) InheritedFromNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) InheritedFromType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) InheritedFromTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inheritedFromTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) InternalValue() *CatalogEffectivePredictiveOptimizationFlag {
	var returns *CatalogEffectivePredictiveOptimizationFlag
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewCatalogEffectivePredictiveOptimizationFlagOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CatalogEffectivePredictiveOptimizationFlagOutputReference {
	_init_.Initialize()

	if err := validateNewCatalogEffectivePredictiveOptimizationFlagOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.catalog.CatalogEffectivePredictiveOptimizationFlagOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCatalogEffectivePredictiveOptimizationFlagOutputReference_Override(c CatalogEffectivePredictiveOptimizationFlagOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.catalog.CatalogEffectivePredictiveOptimizationFlagOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference)SetInheritedFromName(val *string) {
	if err := j.validateSetInheritedFromNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inheritedFromName",
		val,
	)
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference)SetInheritedFromType(val *string) {
	if err := j.validateSetInheritedFromTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inheritedFromType",
		val,
	)
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference)SetInternalValue(val *CatalogEffectivePredictiveOptimizationFlag) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) ResetInheritedFromName() {
	_jsii_.InvokeVoid(
		c,
		"resetInheritedFromName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) ResetInheritedFromType() {
	_jsii_.InvokeVoid(
		c,
		"resetInheritedFromType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEffectivePredictiveOptimizationFlagOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

