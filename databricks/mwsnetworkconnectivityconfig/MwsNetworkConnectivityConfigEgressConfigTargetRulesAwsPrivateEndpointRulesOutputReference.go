// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworkconnectivityconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/mwsnetworkconnectivityconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference interface {
	cdktf.ComplexObject
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
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
	ConnectionState() *string
	SetConnectionState(val *string)
	ConnectionStateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CreationTime() *float64
	SetCreationTime(val *float64)
	CreationTimeInput() *float64
	Deactivated() interface{}
	SetDeactivated(val interface{})
	DeactivatedAt() *float64
	SetDeactivatedAt(val *float64)
	DeactivatedAtInput() *float64
	DeactivatedInput() interface{}
	DomainNames() *[]*string
	SetDomainNames(val *[]*string)
	DomainNamesInput() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EndpointService() *string
	SetEndpointService(val *string)
	EndpointServiceInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NetworkConnectivityConfigId() *string
	SetNetworkConnectivityConfigId(val *string)
	NetworkConnectivityConfigIdInput() *string
	ResourceNames() *[]*string
	SetResourceNames(val *[]*string)
	ResourceNamesInput() *[]*string
	RuleId() *string
	SetRuleId(val *string)
	RuleIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdatedTime() *float64
	SetUpdatedTime(val *float64)
	UpdatedTimeInput() *float64
	VpcEndpointId() *string
	SetVpcEndpointId(val *string)
	VpcEndpointIdInput() *string
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
	ResetAccountId()
	ResetConnectionState()
	ResetCreationTime()
	ResetDeactivated()
	ResetDeactivatedAt()
	ResetDomainNames()
	ResetEnabled()
	ResetEndpointService()
	ResetNetworkConnectivityConfigId()
	ResetResourceNames()
	ResetRuleId()
	ResetUpdatedTime()
	ResetVpcEndpointId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference
type jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ConnectionState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ConnectionStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) CreationTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) Deactivated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) DeactivatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deactivatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) DeactivatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deactivatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) DeactivatedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) DomainNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) DomainNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) EndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) EndpointServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) NetworkConnectivityConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) NetworkConnectivityConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResourceNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResourceNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) RuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) RuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) UpdatedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) UpdatedTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) VpcEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) VpcEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointIdInput",
		&returns,
	)
	return returns
}


func NewMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference {
	_init_.Initialize()

	if err := validateNewMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsNetworkConnectivityConfig.MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference_Override(m MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsNetworkConnectivityConfig.MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetConnectionState(val *string) {
	if err := j.validateSetConnectionStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionState",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetCreationTime(val *float64) {
	if err := j.validateSetCreationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationTime",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetDeactivated(val interface{}) {
	if err := j.validateSetDeactivatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deactivated",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetDeactivatedAt(val *float64) {
	if err := j.validateSetDeactivatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deactivatedAt",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetDomainNames(val *[]*string) {
	if err := j.validateSetDomainNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainNames",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetEndpointService(val *string) {
	if err := j.validateSetEndpointServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointService",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetNetworkConnectivityConfigId(val *string) {
	if err := j.validateSetNetworkConnectivityConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConnectivityConfigId",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetResourceNames(val *[]*string) {
	if err := j.validateSetResourceNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceNames",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetRuleId(val *string) {
	if err := j.validateSetRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleId",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetUpdatedTime(val *float64) {
	if err := j.validateSetUpdatedTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedTime",
		val,
	)
}

func (j *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetVpcEndpointId(val *string) {
	if err := j.validateSetVpcEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEndpointId",
		val,
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetAccountId() {
	_jsii_.InvokeVoid(
		m,
		"resetAccountId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetConnectionState() {
	_jsii_.InvokeVoid(
		m,
		"resetConnectionState",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetCreationTime() {
	_jsii_.InvokeVoid(
		m,
		"resetCreationTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetDeactivated() {
	_jsii_.InvokeVoid(
		m,
		"resetDeactivated",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetDeactivatedAt() {
	_jsii_.InvokeVoid(
		m,
		"resetDeactivatedAt",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetDomainNames() {
	_jsii_.InvokeVoid(
		m,
		"resetDomainNames",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetEndpointService() {
	_jsii_.InvokeVoid(
		m,
		"resetEndpointService",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetNetworkConnectivityConfigId() {
	_jsii_.InvokeVoid(
		m,
		"resetNetworkConnectivityConfigId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetResourceNames() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceNames",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetRuleId() {
	_jsii_.InvokeVoid(
		m,
		"resetRuleId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetUpdatedTime() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdatedTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetVpcEndpointId() {
	_jsii_.InvokeVoid(
		m,
		"resetVpcEndpointId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

