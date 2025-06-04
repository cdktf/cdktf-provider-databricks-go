// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlalert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/sqlalert/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlAlertOptionsOutputReference interface {
	cdktf.ComplexObject
	Column() *string
	SetColumn(val *string)
	ColumnInput() *string
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
	CustomBody() *string
	SetCustomBody(val *string)
	CustomBodyInput() *string
	CustomSubject() *string
	SetCustomSubject(val *string)
	CustomSubjectInput() *string
	EmptyResultState() *string
	SetEmptyResultState(val *string)
	EmptyResultStateInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SqlAlertOptions
	SetInternalValue(val *SqlAlertOptions)
	Muted() interface{}
	SetMuted(val interface{})
	MutedInput() interface{}
	Op() *string
	SetOp(val *string)
	OpInput() *string
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
	ResetCustomBody()
	ResetCustomSubject()
	ResetEmptyResultState()
	ResetMuted()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SqlAlertOptionsOutputReference
type jsiiProxy_SqlAlertOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) Column() *string {
	var returns *string
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) ColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) CustomBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) CustomBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) CustomSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) CustomSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) EmptyResultState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyResultState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) EmptyResultStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyResultStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) InternalValue() *SqlAlertOptions {
	var returns *SqlAlertOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) Muted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"muted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) MutedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) Op() *string {
	var returns *string
	_jsii_.Get(
		j,
		"op",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) OpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewSqlAlertOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SqlAlertOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewSqlAlertOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqlAlertOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.sqlAlert.SqlAlertOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSqlAlertOptionsOutputReference_Override(s SqlAlertOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.sqlAlert.SqlAlertOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetColumn(val *string) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetCustomBody(val *string) {
	if err := j.validateSetCustomBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customBody",
		val,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetCustomSubject(val *string) {
	if err := j.validateSetCustomSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSubject",
		val,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetEmptyResultState(val *string) {
	if err := j.validateSetEmptyResultStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emptyResultState",
		val,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetInternalValue(val *SqlAlertOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetMuted(val interface{}) {
	if err := j.validateSetMutedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"muted",
		val,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetOp(val *string) {
	if err := j.validateSetOpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"op",
		val,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SqlAlertOptionsOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (s *jsiiProxy_SqlAlertOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlAlertOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SqlAlertOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SqlAlertOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SqlAlertOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SqlAlertOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SqlAlertOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SqlAlertOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SqlAlertOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SqlAlertOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SqlAlertOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlAlertOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlAlertOptionsOutputReference) ResetCustomBody() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomBody",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlAlertOptionsOutputReference) ResetCustomSubject() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomSubject",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlAlertOptionsOutputReference) ResetEmptyResultState() {
	_jsii_.InvokeVoid(
		s,
		"resetEmptyResultState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlAlertOptionsOutputReference) ResetMuted() {
	_jsii_.InvokeVoid(
		s,
		"resetMuted",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlAlertOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlAlertOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

