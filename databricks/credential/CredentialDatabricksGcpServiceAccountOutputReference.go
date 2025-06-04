// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package credential

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/credential/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CredentialDatabricksGcpServiceAccountOutputReference interface {
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
	CredentialId() *string
	SetCredentialId(val *string)
	CredentialIdInput() *string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CredentialDatabricksGcpServiceAccount
	SetInternalValue(val *CredentialDatabricksGcpServiceAccount)
	PrivateKeyId() *string
	SetPrivateKeyId(val *string)
	PrivateKeyIdInput() *string
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
	ResetEmail()
	ResetPrivateKeyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CredentialDatabricksGcpServiceAccountOutputReference
type jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) CredentialId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) CredentialIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) InternalValue() *CredentialDatabricksGcpServiceAccount {
	var returns *CredentialDatabricksGcpServiceAccount
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) PrivateKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) PrivateKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCredentialDatabricksGcpServiceAccountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CredentialDatabricksGcpServiceAccountOutputReference {
	_init_.Initialize()

	if err := validateNewCredentialDatabricksGcpServiceAccountOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.credential.CredentialDatabricksGcpServiceAccountOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCredentialDatabricksGcpServiceAccountOutputReference_Override(c CredentialDatabricksGcpServiceAccountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.credential.CredentialDatabricksGcpServiceAccountOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference)SetCredentialId(val *string) {
	if err := j.validateSetCredentialIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialId",
		val,
	)
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference)SetInternalValue(val *CredentialDatabricksGcpServiceAccount) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference)SetPrivateKeyId(val *string) {
	if err := j.validateSetPrivateKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeyId",
		val,
	)
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) ResetCredentialId() {
	_jsii_.InvokeVoid(
		c,
		"resetCredentialId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		c,
		"resetEmail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) ResetPrivateKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivateKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CredentialDatabricksGcpServiceAccountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

