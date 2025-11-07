// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/sqlquery/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlQueryParameterOutputReference interface {
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
	Date() SqlQueryParameterDateOutputReference
	DateInput() *SqlQueryParameterDate
	DateRange() SqlQueryParameterDateRangeOutputReference
	DateRangeInput() *SqlQueryParameterDateRange
	Datetime() SqlQueryParameterDatetimeOutputReference
	DatetimeInput() *SqlQueryParameterDatetime
	DatetimeRange() SqlQueryParameterDatetimeRangeOutputReference
	DatetimeRangeInput() *SqlQueryParameterDatetimeRange
	Datetimesec() SqlQueryParameterDatetimesecOutputReference
	DatetimesecInput() *SqlQueryParameterDatetimesec
	DatetimesecRange() SqlQueryParameterDatetimesecRangeOutputReference
	DatetimesecRangeInput() *SqlQueryParameterDatetimesecRange
	Enum() SqlQueryParameterEnumOutputReference
	EnumInput() *SqlQueryParameterEnum
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Number() SqlQueryParameterNumberOutputReference
	NumberInput() *SqlQueryParameterNumber
	Query() SqlQueryParameterQueryOutputReference
	QueryInput() *SqlQueryParameterQuery
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() SqlQueryParameterTextOutputReference
	TextInput() *SqlQueryParameterText
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDate(value *SqlQueryParameterDate)
	PutDateRange(value *SqlQueryParameterDateRange)
	PutDatetime(value *SqlQueryParameterDatetime)
	PutDatetimeRange(value *SqlQueryParameterDatetimeRange)
	PutDatetimesec(value *SqlQueryParameterDatetimesec)
	PutDatetimesecRange(value *SqlQueryParameterDatetimesecRange)
	PutEnum(value *SqlQueryParameterEnum)
	PutNumber(value *SqlQueryParameterNumber)
	PutQuery(value *SqlQueryParameterQuery)
	PutText(value *SqlQueryParameterText)
	ResetDate()
	ResetDateRange()
	ResetDatetime()
	ResetDatetimeRange()
	ResetDatetimesec()
	ResetDatetimesecRange()
	ResetEnum()
	ResetNumber()
	ResetQuery()
	ResetText()
	ResetTitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SqlQueryParameterOutputReference
type jsiiProxy_SqlQueryParameterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) Date() SqlQueryParameterDateOutputReference {
	var returns SqlQueryParameterDateOutputReference
	_jsii_.Get(
		j,
		"date",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) DateInput() *SqlQueryParameterDate {
	var returns *SqlQueryParameterDate
	_jsii_.Get(
		j,
		"dateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) DateRange() SqlQueryParameterDateRangeOutputReference {
	var returns SqlQueryParameterDateRangeOutputReference
	_jsii_.Get(
		j,
		"dateRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) DateRangeInput() *SqlQueryParameterDateRange {
	var returns *SqlQueryParameterDateRange
	_jsii_.Get(
		j,
		"dateRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) Datetime() SqlQueryParameterDatetimeOutputReference {
	var returns SqlQueryParameterDatetimeOutputReference
	_jsii_.Get(
		j,
		"datetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) DatetimeInput() *SqlQueryParameterDatetime {
	var returns *SqlQueryParameterDatetime
	_jsii_.Get(
		j,
		"datetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) DatetimeRange() SqlQueryParameterDatetimeRangeOutputReference {
	var returns SqlQueryParameterDatetimeRangeOutputReference
	_jsii_.Get(
		j,
		"datetimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) DatetimeRangeInput() *SqlQueryParameterDatetimeRange {
	var returns *SqlQueryParameterDatetimeRange
	_jsii_.Get(
		j,
		"datetimeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) Datetimesec() SqlQueryParameterDatetimesecOutputReference {
	var returns SqlQueryParameterDatetimesecOutputReference
	_jsii_.Get(
		j,
		"datetimesec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) DatetimesecInput() *SqlQueryParameterDatetimesec {
	var returns *SqlQueryParameterDatetimesec
	_jsii_.Get(
		j,
		"datetimesecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) DatetimesecRange() SqlQueryParameterDatetimesecRangeOutputReference {
	var returns SqlQueryParameterDatetimesecRangeOutputReference
	_jsii_.Get(
		j,
		"datetimesecRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) DatetimesecRangeInput() *SqlQueryParameterDatetimesecRange {
	var returns *SqlQueryParameterDatetimesecRange
	_jsii_.Get(
		j,
		"datetimesecRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) Enum() SqlQueryParameterEnumOutputReference {
	var returns SqlQueryParameterEnumOutputReference
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) EnumInput() *SqlQueryParameterEnum {
	var returns *SqlQueryParameterEnum
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) Number() SqlQueryParameterNumberOutputReference {
	var returns SqlQueryParameterNumberOutputReference
	_jsii_.Get(
		j,
		"number",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) NumberInput() *SqlQueryParameterNumber {
	var returns *SqlQueryParameterNumber
	_jsii_.Get(
		j,
		"numberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) Query() SqlQueryParameterQueryOutputReference {
	var returns SqlQueryParameterQueryOutputReference
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) QueryInput() *SqlQueryParameterQuery {
	var returns *SqlQueryParameterQuery
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) Text() SqlQueryParameterTextOutputReference {
	var returns SqlQueryParameterTextOutputReference
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) TextInput() *SqlQueryParameterText {
	var returns *SqlQueryParameterText
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlQueryParameterOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


func NewSqlQueryParameterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SqlQueryParameterOutputReference {
	_init_.Initialize()

	if err := validateNewSqlQueryParameterOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqlQueryParameterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.sqlQuery.SqlQueryParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSqlQueryParameterOutputReference_Override(s SqlQueryParameterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.sqlQuery.SqlQueryParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SqlQueryParameterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SqlQueryParameterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SqlQueryParameterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SqlQueryParameterOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SqlQueryParameterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SqlQueryParameterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SqlQueryParameterOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) PutDate(value *SqlQueryParameterDate) {
	if err := s.validatePutDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) PutDateRange(value *SqlQueryParameterDateRange) {
	if err := s.validatePutDateRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDateRange",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) PutDatetime(value *SqlQueryParameterDatetime) {
	if err := s.validatePutDatetimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatetime",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) PutDatetimeRange(value *SqlQueryParameterDatetimeRange) {
	if err := s.validatePutDatetimeRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatetimeRange",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) PutDatetimesec(value *SqlQueryParameterDatetimesec) {
	if err := s.validatePutDatetimesecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatetimesec",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) PutDatetimesecRange(value *SqlQueryParameterDatetimesecRange) {
	if err := s.validatePutDatetimesecRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatetimesecRange",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) PutEnum(value *SqlQueryParameterEnum) {
	if err := s.validatePutEnumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEnum",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) PutNumber(value *SqlQueryParameterNumber) {
	if err := s.validatePutNumberParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNumber",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) PutQuery(value *SqlQueryParameterQuery) {
	if err := s.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putQuery",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) PutText(value *SqlQueryParameterText) {
	if err := s.validatePutTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putText",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ResetDate() {
	_jsii_.InvokeVoid(
		s,
		"resetDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ResetDateRange() {
	_jsii_.InvokeVoid(
		s,
		"resetDateRange",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ResetDatetime() {
	_jsii_.InvokeVoid(
		s,
		"resetDatetime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ResetDatetimeRange() {
	_jsii_.InvokeVoid(
		s,
		"resetDatetimeRange",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ResetDatetimesec() {
	_jsii_.InvokeVoid(
		s,
		"resetDatetimesec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ResetDatetimesecRange() {
	_jsii_.InvokeVoid(
		s,
		"resetDatetimesecRange",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ResetEnum() {
	_jsii_.InvokeVoid(
		s,
		"resetEnum",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ResetNumber() {
	_jsii_.InvokeVoid(
		s,
		"resetNumber",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		s,
		"resetQuery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		s,
		"resetText",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		s,
		"resetTitle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlQueryParameterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

