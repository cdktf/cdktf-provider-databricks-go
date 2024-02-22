// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksstoragecredential

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabricksstoragecredential/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference interface {
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
	InternalValue() *DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity
	SetInternalValue(val *DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity)
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

// The jsii proxy struct for DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference
type jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) AccessConnectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessConnectorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) AccessConnectorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessConnectorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) CredentialId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) CredentialIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) InternalValue() *DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity {
	var returns *DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) ManagedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) ManagedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksStorageCredential.DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference_Override(d DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksStorageCredential.DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference)SetAccessConnectorId(val *string) {
	if err := j.validateSetAccessConnectorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessConnectorId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference)SetCredentialId(val *string) {
	if err := j.validateSetCredentialIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference)SetInternalValue(val *DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference)SetManagedIdentityId(val *string) {
	if err := j.validateSetManagedIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedIdentityId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) ResetCredentialId() {
	_jsii_.InvokeVoid(
		d,
		"resetCredentialId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) ResetManagedIdentityId() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedIdentityId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

