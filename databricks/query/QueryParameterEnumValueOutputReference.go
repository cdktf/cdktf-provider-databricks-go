// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/query/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QueryParameterEnumValueOutputReference interface {
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
	EnumOptions() *string
	SetEnumOptions(val *string)
	EnumOptionsInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *QueryParameterEnumValue
	SetInternalValue(val *QueryParameterEnumValue)
	MultiValuesOptions() QueryParameterEnumValueMultiValuesOptionsOutputReference
	MultiValuesOptionsInput() *QueryParameterEnumValueMultiValuesOptions
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	PutMultiValuesOptions(value *QueryParameterEnumValueMultiValuesOptions)
	ResetEnumOptions()
	ResetMultiValuesOptions()
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QueryParameterEnumValueOutputReference
type jsiiProxy_QueryParameterEnumValueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) EnumOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enumOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) EnumOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enumOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) InternalValue() *QueryParameterEnumValue {
	var returns *QueryParameterEnumValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) MultiValuesOptions() QueryParameterEnumValueMultiValuesOptionsOutputReference {
	var returns QueryParameterEnumValueMultiValuesOptionsOutputReference
	_jsii_.Get(
		j,
		"multiValuesOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) MultiValuesOptionsInput() *QueryParameterEnumValueMultiValuesOptions {
	var returns *QueryParameterEnumValueMultiValuesOptions
	_jsii_.Get(
		j,
		"multiValuesOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewQueryParameterEnumValueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QueryParameterEnumValueOutputReference {
	_init_.Initialize()

	if err := validateNewQueryParameterEnumValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QueryParameterEnumValueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.query.QueryParameterEnumValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQueryParameterEnumValueOutputReference_Override(q QueryParameterEnumValueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.query.QueryParameterEnumValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference)SetEnumOptions(val *string) {
	if err := j.validateSetEnumOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enumOptions",
		val,
	)
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference)SetInternalValue(val *QueryParameterEnumValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QueryParameterEnumValueOutputReference)SetValues(val *[]*string) {
	if err := j.validateSetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) PutMultiValuesOptions(value *QueryParameterEnumValueMultiValuesOptions) {
	if err := q.validatePutMultiValuesOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMultiValuesOptions",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) ResetEnumOptions() {
	_jsii_.InvokeVoid(
		q,
		"resetEnumOptions",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) ResetMultiValuesOptions() {
	_jsii_.InvokeVoid(
		q,
		"resetMultiValuesOptions",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		q,
		"resetValues",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterEnumValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

