// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworkconnectivityconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/mwsnetworkconnectivityconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference interface {
	cdktf.ComplexObject
	AwsStableIpRule() MwsNetworkConnectivityConfigEgressConfigDefaultRulesAwsStableIpRuleOutputReference
	AwsStableIpRuleInput() *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAwsStableIpRule
	AzureServiceEndpointRule() MwsNetworkConnectivityConfigEgressConfigDefaultRulesAzureServiceEndpointRuleOutputReference
	AzureServiceEndpointRuleInput() *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAzureServiceEndpointRule
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
	InternalValue() *MwsNetworkConnectivityConfigEgressConfigDefaultRules
	SetInternalValue(val *MwsNetworkConnectivityConfigEgressConfigDefaultRules)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAwsStableIpRule(value *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAwsStableIpRule)
	PutAzureServiceEndpointRule(value *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAzureServiceEndpointRule)
	ResetAwsStableIpRule()
	ResetAzureServiceEndpointRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference
type jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) AwsStableIpRule() MwsNetworkConnectivityConfigEgressConfigDefaultRulesAwsStableIpRuleOutputReference {
	var returns MwsNetworkConnectivityConfigEgressConfigDefaultRulesAwsStableIpRuleOutputReference
	_jsii_.Get(
		j,
		"awsStableIpRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) AwsStableIpRuleInput() *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAwsStableIpRule {
	var returns *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAwsStableIpRule
	_jsii_.Get(
		j,
		"awsStableIpRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) AzureServiceEndpointRule() MwsNetworkConnectivityConfigEgressConfigDefaultRulesAzureServiceEndpointRuleOutputReference {
	var returns MwsNetworkConnectivityConfigEgressConfigDefaultRulesAzureServiceEndpointRuleOutputReference
	_jsii_.Get(
		j,
		"azureServiceEndpointRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) AzureServiceEndpointRuleInput() *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAzureServiceEndpointRule {
	var returns *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAzureServiceEndpointRule
	_jsii_.Get(
		j,
		"azureServiceEndpointRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) InternalValue() *MwsNetworkConnectivityConfigEgressConfigDefaultRules {
	var returns *MwsNetworkConnectivityConfigEgressConfigDefaultRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference {
	_init_.Initialize()

	if err := validateNewMwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsNetworkConnectivityConfig.MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference_Override(m MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsNetworkConnectivityConfig.MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference)SetInternalValue(val *MwsNetworkConnectivityConfigEgressConfigDefaultRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) PutAwsStableIpRule(value *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAwsStableIpRule) {
	if err := m.validatePutAwsStableIpRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAwsStableIpRule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) PutAzureServiceEndpointRule(value *MwsNetworkConnectivityConfigEgressConfigDefaultRulesAzureServiceEndpointRule) {
	if err := m.validatePutAzureServiceEndpointRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAzureServiceEndpointRule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) ResetAwsStableIpRule() {
	_jsii_.InvokeVoid(
		m,
		"resetAwsStableIpRule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) ResetAzureServiceEndpointRule() {
	_jsii_.InvokeVoid(
		m,
		"resetAzureServiceEndpointRule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigDefaultRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

