// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/query/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QueryParameterDateRangeValueOutputReference interface {
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
	DateRangeValue() QueryParameterDateRangeValueDateRangeValueOutputReference
	DateRangeValueInput() *QueryParameterDateRangeValueDateRangeValue
	DynamicDateRangeValue() *string
	SetDynamicDateRangeValue(val *string)
	DynamicDateRangeValueInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *QueryParameterDateRangeValue
	SetInternalValue(val *QueryParameterDateRangeValue)
	Precision() *string
	SetPrecision(val *string)
	PrecisionInput() *string
	StartDayOfWeek() *float64
	SetStartDayOfWeek(val *float64)
	StartDayOfWeekInput() *float64
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
	PutDateRangeValue(value *QueryParameterDateRangeValueDateRangeValue)
	ResetDateRangeValue()
	ResetDynamicDateRangeValue()
	ResetPrecision()
	ResetStartDayOfWeek()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QueryParameterDateRangeValueOutputReference
type jsiiProxy_QueryParameterDateRangeValueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) DateRangeValue() QueryParameterDateRangeValueDateRangeValueOutputReference {
	var returns QueryParameterDateRangeValueDateRangeValueOutputReference
	_jsii_.Get(
		j,
		"dateRangeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) DateRangeValueInput() *QueryParameterDateRangeValueDateRangeValue {
	var returns *QueryParameterDateRangeValueDateRangeValue
	_jsii_.Get(
		j,
		"dateRangeValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) DynamicDateRangeValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynamicDateRangeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) DynamicDateRangeValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynamicDateRangeValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) InternalValue() *QueryParameterDateRangeValue {
	var returns *QueryParameterDateRangeValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) Precision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"precision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) PrecisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"precisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) StartDayOfWeek() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startDayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) StartDayOfWeekInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startDayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQueryParameterDateRangeValueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QueryParameterDateRangeValueOutputReference {
	_init_.Initialize()

	if err := validateNewQueryParameterDateRangeValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QueryParameterDateRangeValueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.query.QueryParameterDateRangeValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQueryParameterDateRangeValueOutputReference_Override(q QueryParameterDateRangeValueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.query.QueryParameterDateRangeValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference)SetDynamicDateRangeValue(val *string) {
	if err := j.validateSetDynamicDateRangeValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicDateRangeValue",
		val,
	)
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference)SetInternalValue(val *QueryParameterDateRangeValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference)SetPrecision(val *string) {
	if err := j.validateSetPrecisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precision",
		val,
	)
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference)SetStartDayOfWeek(val *float64) {
	if err := j.validateSetStartDayOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startDayOfWeek",
		val,
	)
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QueryParameterDateRangeValueOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) PutDateRangeValue(value *QueryParameterDateRangeValueDateRangeValue) {
	if err := q.validatePutDateRangeValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDateRangeValue",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) ResetDateRangeValue() {
	_jsii_.InvokeVoid(
		q,
		"resetDateRangeValue",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) ResetDynamicDateRangeValue() {
	_jsii_.InvokeVoid(
		q,
		"resetDynamicDateRangeValue",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) ResetPrecision() {
	_jsii_.InvokeVoid(
		q,
		"resetPrecision",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) ResetStartDayOfWeek() {
	_jsii_.InvokeVoid(
		q,
		"resetStartDayOfWeek",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (q *jsiiProxy_QueryParameterDateRangeValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

