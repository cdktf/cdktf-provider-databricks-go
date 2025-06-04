// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/accountnetworkpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountNetworkPolicyEgressNetworkAccessOutputReference interface {
	cdktf.ComplexObject
	AllowedInternetDestinations() AccountNetworkPolicyEgressNetworkAccessAllowedInternetDestinationsList
	AllowedInternetDestinationsInput() interface{}
	AllowedStorageDestinations() AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsList
	AllowedStorageDestinationsInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PolicyEnforcement() AccountNetworkPolicyEgressNetworkAccessPolicyEnforcementOutputReference
	PolicyEnforcementInput() interface{}
	RestrictionMode() *string
	SetRestrictionMode(val *string)
	RestrictionModeInput() *string
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
	PutAllowedInternetDestinations(value interface{})
	PutAllowedStorageDestinations(value interface{})
	PutPolicyEnforcement(value *AccountNetworkPolicyEgressNetworkAccessPolicyEnforcement)
	ResetAllowedInternetDestinations()
	ResetAllowedStorageDestinations()
	ResetPolicyEnforcement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountNetworkPolicyEgressNetworkAccessOutputReference
type jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) AllowedInternetDestinations() AccountNetworkPolicyEgressNetworkAccessAllowedInternetDestinationsList {
	var returns AccountNetworkPolicyEgressNetworkAccessAllowedInternetDestinationsList
	_jsii_.Get(
		j,
		"allowedInternetDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) AllowedInternetDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedInternetDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) AllowedStorageDestinations() AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsList {
	var returns AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinationsList
	_jsii_.Get(
		j,
		"allowedStorageDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) AllowedStorageDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedStorageDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) PolicyEnforcement() AccountNetworkPolicyEgressNetworkAccessPolicyEnforcementOutputReference {
	var returns AccountNetworkPolicyEgressNetworkAccessPolicyEnforcementOutputReference
	_jsii_.Get(
		j,
		"policyEnforcement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) PolicyEnforcementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyEnforcementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) RestrictionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) RestrictionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountNetworkPolicyEgressNetworkAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccountNetworkPolicyEgressNetworkAccessOutputReference {
	_init_.Initialize()

	if err := validateNewAccountNetworkPolicyEgressNetworkAccessOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyEgressNetworkAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountNetworkPolicyEgressNetworkAccessOutputReference_Override(a AccountNetworkPolicyEgressNetworkAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyEgressNetworkAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference)SetRestrictionMode(val *string) {
	if err := j.validateSetRestrictionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictionMode",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) PutAllowedInternetDestinations(value interface{}) {
	if err := a.validatePutAllowedInternetDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAllowedInternetDestinations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) PutAllowedStorageDestinations(value interface{}) {
	if err := a.validatePutAllowedStorageDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAllowedStorageDestinations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) PutPolicyEnforcement(value *AccountNetworkPolicyEgressNetworkAccessPolicyEnforcement) {
	if err := a.validatePutPolicyEnforcementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPolicyEnforcement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) ResetAllowedInternetDestinations() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedInternetDestinations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) ResetAllowedStorageDestinations() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedStorageDestinations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) ResetPolicyEnforcement() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyEnforcement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyEgressNetworkAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

