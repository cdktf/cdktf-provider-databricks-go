// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/query/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QueryParameterOutputReference interface {
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
	DateRangeValue() QueryParameterDateRangeValueOutputReference
	DateRangeValueInput() *QueryParameterDateRangeValue
	DateValue() QueryParameterDateValueOutputReference
	DateValueInput() *QueryParameterDateValue
	EnumValue() QueryParameterEnumValueOutputReference
	EnumValueInput() *QueryParameterEnumValue
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	NumericValue() QueryParameterNumericValueOutputReference
	NumericValueInput() *QueryParameterNumericValue
	QueryBackedValue() QueryParameterQueryBackedValueOutputReference
	QueryBackedValueInput() *QueryParameterQueryBackedValue
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextValue() QueryParameterTextValueOutputReference
	TextValueInput() *QueryParameterTextValue
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
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
	PutDateRangeValue(value *QueryParameterDateRangeValue)
	PutDateValue(value *QueryParameterDateValue)
	PutEnumValue(value *QueryParameterEnumValue)
	PutNumericValue(value *QueryParameterNumericValue)
	PutQueryBackedValue(value *QueryParameterQueryBackedValue)
	PutTextValue(value *QueryParameterTextValue)
	ResetDateRangeValue()
	ResetDateValue()
	ResetEnumValue()
	ResetNumericValue()
	ResetQueryBackedValue()
	ResetTextValue()
	ResetTitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QueryParameterOutputReference
type jsiiProxy_QueryParameterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QueryParameterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) DateRangeValue() QueryParameterDateRangeValueOutputReference {
	var returns QueryParameterDateRangeValueOutputReference
	_jsii_.Get(
		j,
		"dateRangeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) DateRangeValueInput() *QueryParameterDateRangeValue {
	var returns *QueryParameterDateRangeValue
	_jsii_.Get(
		j,
		"dateRangeValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) DateValue() QueryParameterDateValueOutputReference {
	var returns QueryParameterDateValueOutputReference
	_jsii_.Get(
		j,
		"dateValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) DateValueInput() *QueryParameterDateValue {
	var returns *QueryParameterDateValue
	_jsii_.Get(
		j,
		"dateValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) EnumValue() QueryParameterEnumValueOutputReference {
	var returns QueryParameterEnumValueOutputReference
	_jsii_.Get(
		j,
		"enumValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) EnumValueInput() *QueryParameterEnumValue {
	var returns *QueryParameterEnumValue
	_jsii_.Get(
		j,
		"enumValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) NumericValue() QueryParameterNumericValueOutputReference {
	var returns QueryParameterNumericValueOutputReference
	_jsii_.Get(
		j,
		"numericValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) NumericValueInput() *QueryParameterNumericValue {
	var returns *QueryParameterNumericValue
	_jsii_.Get(
		j,
		"numericValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) QueryBackedValue() QueryParameterQueryBackedValueOutputReference {
	var returns QueryParameterQueryBackedValueOutputReference
	_jsii_.Get(
		j,
		"queryBackedValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) QueryBackedValueInput() *QueryParameterQueryBackedValue {
	var returns *QueryParameterQueryBackedValue
	_jsii_.Get(
		j,
		"queryBackedValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) TextValue() QueryParameterTextValueOutputReference {
	var returns QueryParameterTextValueOutputReference
	_jsii_.Get(
		j,
		"textValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) TextValueInput() *QueryParameterTextValue {
	var returns *QueryParameterTextValue
	_jsii_.Get(
		j,
		"textValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryParameterOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


func NewQueryParameterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QueryParameterOutputReference {
	_init_.Initialize()

	if err := validateNewQueryParameterOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QueryParameterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.query.QueryParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQueryParameterOutputReference_Override(q QueryParameterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.query.QueryParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QueryParameterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QueryParameterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QueryParameterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QueryParameterOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_QueryParameterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QueryParameterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QueryParameterOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QueryParameterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QueryParameterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QueryParameterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QueryParameterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QueryParameterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QueryParameterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QueryParameterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QueryParameterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QueryParameterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QueryParameterOutputReference) PutDateRangeValue(value *QueryParameterDateRangeValue) {
	if err := q.validatePutDateRangeValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDateRangeValue",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) PutDateValue(value *QueryParameterDateValue) {
	if err := q.validatePutDateValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDateValue",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) PutEnumValue(value *QueryParameterEnumValue) {
	if err := q.validatePutEnumValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putEnumValue",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) PutNumericValue(value *QueryParameterNumericValue) {
	if err := q.validatePutNumericValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putNumericValue",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) PutQueryBackedValue(value *QueryParameterQueryBackedValue) {
	if err := q.validatePutQueryBackedValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putQueryBackedValue",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) PutTextValue(value *QueryParameterTextValue) {
	if err := q.validatePutTextValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTextValue",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) ResetDateRangeValue() {
	_jsii_.InvokeVoid(
		q,
		"resetDateRangeValue",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) ResetDateValue() {
	_jsii_.InvokeVoid(
		q,
		"resetDateValue",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) ResetEnumValue() {
	_jsii_.InvokeVoid(
		q,
		"resetEnumValue",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) ResetNumericValue() {
	_jsii_.InvokeVoid(
		q,
		"resetNumericValue",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) ResetQueryBackedValue() {
	_jsii_.InvokeVoid(
		q,
		"resetQueryBackedValue",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) ResetTextValue() {
	_jsii_.InvokeVoid(
		q,
		"resetTextValue",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		q,
		"resetTitle",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryParameterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (q *jsiiProxy_QueryParameterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

