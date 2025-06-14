// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metastoredataaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/metastoredataaccess/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MetastoreDataAccessAzureManagedIdentityOutputReference interface {
	cdktf.ComplexObject
	AccessConnectorId() *string
	SetAccessConnectorId(val *string)
	AccessConnectorIdInput() *string
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
	CredentialId() *string
	SetCredentialId(val *string)
	CredentialIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MetastoreDataAccessAzureManagedIdentity
	SetInternalValue(val *MetastoreDataAccessAzureManagedIdentity)
	ManagedIdentityId() *string
	SetManagedIdentityId(val *string)
	ManagedIdentityIdInput() *string
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
	ResetCredentialId()
	ResetManagedIdentityId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MetastoreDataAccessAzureManagedIdentityOutputReference
type jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) AccessConnectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessConnectorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) AccessConnectorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessConnectorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) CredentialId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) CredentialIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) InternalValue() *MetastoreDataAccessAzureManagedIdentity {
	var returns *MetastoreDataAccessAzureManagedIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) ManagedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) ManagedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMetastoreDataAccessAzureManagedIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MetastoreDataAccessAzureManagedIdentityOutputReference {
	_init_.Initialize()

	if err := validateNewMetastoreDataAccessAzureManagedIdentityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.metastoreDataAccess.MetastoreDataAccessAzureManagedIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMetastoreDataAccessAzureManagedIdentityOutputReference_Override(m MetastoreDataAccessAzureManagedIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.metastoreDataAccess.MetastoreDataAccessAzureManagedIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference)SetAccessConnectorId(val *string) {
	if err := j.validateSetAccessConnectorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessConnectorId",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference)SetCredentialId(val *string) {
	if err := j.validateSetCredentialIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialId",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference)SetInternalValue(val *MetastoreDataAccessAzureManagedIdentity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference)SetManagedIdentityId(val *string) {
	if err := j.validateSetManagedIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedIdentityId",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) ResetCredentialId() {
	_jsii_.InvokeVoid(
		m,
		"resetCredentialId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) ResetManagedIdentityId() {
	_jsii_.InvokeVoid(
		m,
		"resetManagedIdentityId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MetastoreDataAccessAzureManagedIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

