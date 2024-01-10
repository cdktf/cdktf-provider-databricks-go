// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwscustomermanagedkeys

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/mwscustomermanagedkeys/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsCustomerManagedKeysAwsKeyInfoOutputReference interface {
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
	InternalValue() *MwsCustomerManagedKeysAwsKeyInfo
	SetInternalValue(val *MwsCustomerManagedKeysAwsKeyInfo)
	KeyAlias() *string
	SetKeyAlias(val *string)
	KeyAliasInput() *string
	KeyArn() *string
	SetKeyArn(val *string)
	KeyArnInput() *string
	KeyRegion() *string
	SetKeyRegion(val *string)
	KeyRegionInput() *string
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
	ResetKeyAlias()
	ResetKeyRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MwsCustomerManagedKeysAwsKeyInfoOutputReference
type jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) InternalValue() *MwsCustomerManagedKeysAwsKeyInfo {
	var returns *MwsCustomerManagedKeysAwsKeyInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) KeyAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) KeyAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) KeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) KeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) KeyRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) KeyRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMwsCustomerManagedKeysAwsKeyInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MwsCustomerManagedKeysAwsKeyInfoOutputReference {
	_init_.Initialize()

	if err := validateNewMwsCustomerManagedKeysAwsKeyInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsCustomerManagedKeys.MwsCustomerManagedKeysAwsKeyInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMwsCustomerManagedKeysAwsKeyInfoOutputReference_Override(m MwsCustomerManagedKeysAwsKeyInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsCustomerManagedKeys.MwsCustomerManagedKeysAwsKeyInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference)SetInternalValue(val *MwsCustomerManagedKeysAwsKeyInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference)SetKeyAlias(val *string) {
	if err := j.validateSetKeyAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyAlias",
		val,
	)
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference)SetKeyArn(val *string) {
	if err := j.validateSetKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyArn",
		val,
	)
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference)SetKeyRegion(val *string) {
	if err := j.validateSetKeyRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyRegion",
		val,
	)
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) ResetKeyAlias() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyAlias",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) ResetKeyRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MwsCustomerManagedKeysAwsKeyInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

