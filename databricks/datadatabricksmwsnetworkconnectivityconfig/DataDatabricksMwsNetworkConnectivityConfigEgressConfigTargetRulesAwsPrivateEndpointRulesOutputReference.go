// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmwsnetworkconnectivityconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksmwsnetworkconnectivityconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference interface {
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
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
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference
type jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ConnectionState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ConnectionStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) CreationTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) Deactivated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) DeactivatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deactivatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) DeactivatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deactivatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) DeactivatedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) DomainNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) DomainNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) EndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) EndpointServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) NetworkConnectivityConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) NetworkConnectivityConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResourceNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResourceNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) RuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) RuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) UpdatedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) UpdatedTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) VpcEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) VpcEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointIdInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksMwsNetworkConnectivityConfig.DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference_Override(d DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksMwsNetworkConnectivityConfig.DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetConnectionState(val *string) {
	if err := j.validateSetConnectionStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionState",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetCreationTime(val *float64) {
	if err := j.validateSetCreationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetDeactivated(val interface{}) {
	if err := j.validateSetDeactivatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deactivated",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetDeactivatedAt(val *float64) {
	if err := j.validateSetDeactivatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deactivatedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetDomainNames(val *[]*string) {
	if err := j.validateSetDomainNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainNames",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetEndpointService(val *string) {
	if err := j.validateSetEndpointServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointService",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetNetworkConnectivityConfigId(val *string) {
	if err := j.validateSetNetworkConnectivityConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConnectivityConfigId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetResourceNames(val *[]*string) {
	if err := j.validateSetResourceNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceNames",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetRuleId(val *string) {
	if err := j.validateSetRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetUpdatedTime(val *float64) {
	if err := j.validateSetUpdatedTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference)SetVpcEndpointId(val *string) {
	if err := j.validateSetVpcEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEndpointId",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetAccountId() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetConnectionState() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetCreationTime() {
	_jsii_.InvokeVoid(
		d,
		"resetCreationTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetDeactivated() {
	_jsii_.InvokeVoid(
		d,
		"resetDeactivated",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetDeactivatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetDeactivatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetDomainNames() {
	_jsii_.InvokeVoid(
		d,
		"resetDomainNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetEndpointService() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetNetworkConnectivityConfigId() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkConnectivityConfigId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetResourceNames() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetRuleId() {
	_jsii_.InvokeVoid(
		d,
		"resetRuleId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetUpdatedTime() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ResetVpcEndpointId() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcEndpointId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAwsPrivateEndpointRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

