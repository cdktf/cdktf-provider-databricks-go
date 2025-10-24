// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customappintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/customappintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomAppIntegrationTokenAccessPolicyOutputReference interface {
	cdktf.ComplexObject
	AbsoluteSessionLifetimeInMinutes() *float64
	SetAbsoluteSessionLifetimeInMinutes(val *float64)
	AbsoluteSessionLifetimeInMinutesInput() *float64
	AccessTokenTtlInMinutes() *float64
	SetAccessTokenTtlInMinutes(val *float64)
	AccessTokenTtlInMinutesInput() *float64
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
	EnableSingleUseRefreshTokens() interface{}
	SetEnableSingleUseRefreshTokens(val interface{})
	EnableSingleUseRefreshTokensInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CustomAppIntegrationTokenAccessPolicy
	SetInternalValue(val *CustomAppIntegrationTokenAccessPolicy)
	RefreshTokenTtlInMinutes() *float64
	SetRefreshTokenTtlInMinutes(val *float64)
	RefreshTokenTtlInMinutesInput() *float64
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
	ResetAbsoluteSessionLifetimeInMinutes()
	ResetAccessTokenTtlInMinutes()
	ResetEnableSingleUseRefreshTokens()
	ResetRefreshTokenTtlInMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomAppIntegrationTokenAccessPolicyOutputReference
type jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) AbsoluteSessionLifetimeInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"absoluteSessionLifetimeInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) AbsoluteSessionLifetimeInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"absoluteSessionLifetimeInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) AccessTokenTtlInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenTtlInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) AccessTokenTtlInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenTtlInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) EnableSingleUseRefreshTokens() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSingleUseRefreshTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) EnableSingleUseRefreshTokensInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSingleUseRefreshTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) InternalValue() *CustomAppIntegrationTokenAccessPolicy {
	var returns *CustomAppIntegrationTokenAccessPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) RefreshTokenTtlInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenTtlInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) RefreshTokenTtlInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenTtlInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCustomAppIntegrationTokenAccessPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomAppIntegrationTokenAccessPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewCustomAppIntegrationTokenAccessPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.customAppIntegration.CustomAppIntegrationTokenAccessPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomAppIntegrationTokenAccessPolicyOutputReference_Override(c CustomAppIntegrationTokenAccessPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.customAppIntegration.CustomAppIntegrationTokenAccessPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference)SetAbsoluteSessionLifetimeInMinutes(val *float64) {
	if err := j.validateSetAbsoluteSessionLifetimeInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"absoluteSessionLifetimeInMinutes",
		val,
	)
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference)SetAccessTokenTtlInMinutes(val *float64) {
	if err := j.validateSetAccessTokenTtlInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessTokenTtlInMinutes",
		val,
	)
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference)SetEnableSingleUseRefreshTokens(val interface{}) {
	if err := j.validateSetEnableSingleUseRefreshTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSingleUseRefreshTokens",
		val,
	)
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference)SetInternalValue(val *CustomAppIntegrationTokenAccessPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference)SetRefreshTokenTtlInMinutes(val *float64) {
	if err := j.validateSetRefreshTokenTtlInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTokenTtlInMinutes",
		val,
	)
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) ResetAbsoluteSessionLifetimeInMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetAbsoluteSessionLifetimeInMinutes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) ResetAccessTokenTtlInMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessTokenTtlInMinutes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) ResetEnableSingleUseRefreshTokens() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableSingleUseRefreshTokens",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) ResetRefreshTokenTtlInMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetRefreshTokenTtlInMinutes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CustomAppIntegrationTokenAccessPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

