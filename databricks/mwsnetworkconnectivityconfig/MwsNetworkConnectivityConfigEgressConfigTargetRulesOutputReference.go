// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworkconnectivityconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/mwsnetworkconnectivityconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference interface {
	cdktf.ComplexObject
	AwsPrivateEndpointRules() MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesList
	AwsPrivateEndpointRulesInput() interface{}
	AzurePrivateEndpointRules() MwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesList
	AzurePrivateEndpointRulesInput() interface{}
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
	InternalValue() *MwsNetworkConnectivityConfigEgressConfigTargetRules
	SetInternalValue(val *MwsNetworkConnectivityConfigEgressConfigTargetRules)
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
	PutAwsPrivateEndpointRules(value interface{})
	PutAzurePrivateEndpointRules(value interface{})
	ResetAwsPrivateEndpointRules()
	ResetAzurePrivateEndpointRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference
type jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) AwsPrivateEndpointRules() MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesList {
	var returns MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesList
	_jsii_.Get(
		j,
		"awsPrivateEndpointRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) AwsPrivateEndpointRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsPrivateEndpointRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) AzurePrivateEndpointRules() MwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesList {
	var returns MwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesList
	_jsii_.Get(
		j,
		"azurePrivateEndpointRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) AzurePrivateEndpointRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azurePrivateEndpointRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) InternalValue() *MwsNetworkConnectivityConfigEgressConfigTargetRules {
	var returns *MwsNetworkConnectivityConfigEgressConfigTargetRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference {
	_init_.Initialize()

	if err := validateNewMwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsNetworkConnectivityConfig.MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference_Override(m MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsNetworkConnectivityConfig.MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference)SetInternalValue(val *MwsNetworkConnectivityConfigEgressConfigTargetRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) PutAwsPrivateEndpointRules(value interface{}) {
	if err := m.validatePutAwsPrivateEndpointRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAwsPrivateEndpointRules",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) PutAzurePrivateEndpointRules(value interface{}) {
	if err := m.validatePutAzurePrivateEndpointRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAzurePrivateEndpointRules",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) ResetAwsPrivateEndpointRules() {
	_jsii_.InvokeVoid(
		m,
		"resetAwsPrivateEndpointRules",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) ResetAzurePrivateEndpointRules() {
	_jsii_.InvokeVoid(
		m,
		"resetAzurePrivateEndpointRules",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

