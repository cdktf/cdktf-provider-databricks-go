// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsworkspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v10/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v10/mwsworkspaces/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsWorkspacesExternalCustomerInfoOutputReference interface {
	cdktf.ComplexObject
	AuthoritativeUserEmail() *string
	SetAuthoritativeUserEmail(val *string)
	AuthoritativeUserEmailInput() *string
	AuthoritativeUserFullName() *string
	SetAuthoritativeUserFullName(val *string)
	AuthoritativeUserFullNameInput() *string
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
	CustomerName() *string
	SetCustomerName(val *string)
	CustomerNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MwsWorkspacesExternalCustomerInfo
	SetInternalValue(val *MwsWorkspacesExternalCustomerInfo)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MwsWorkspacesExternalCustomerInfoOutputReference
type jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) AuthoritativeUserEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authoritativeUserEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) AuthoritativeUserEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authoritativeUserEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) AuthoritativeUserFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authoritativeUserFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) AuthoritativeUserFullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authoritativeUserFullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) CustomerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) CustomerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) InternalValue() *MwsWorkspacesExternalCustomerInfo {
	var returns *MwsWorkspacesExternalCustomerInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMwsWorkspacesExternalCustomerInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MwsWorkspacesExternalCustomerInfoOutputReference {
	_init_.Initialize()

	if err := validateNewMwsWorkspacesExternalCustomerInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsWorkspaces.MwsWorkspacesExternalCustomerInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMwsWorkspacesExternalCustomerInfoOutputReference_Override(m MwsWorkspacesExternalCustomerInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsWorkspaces.MwsWorkspacesExternalCustomerInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference)SetAuthoritativeUserEmail(val *string) {
	if err := j.validateSetAuthoritativeUserEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authoritativeUserEmail",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference)SetAuthoritativeUserFullName(val *string) {
	if err := j.validateSetAuthoritativeUserFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authoritativeUserFullName",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference)SetCustomerName(val *string) {
	if err := j.validateSetCustomerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerName",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference)SetInternalValue(val *MwsWorkspacesExternalCustomerInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesExternalCustomerInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

