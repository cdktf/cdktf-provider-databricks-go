// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmwsnetworkconnectivityconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksmwsnetworkconnectivityconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference interface {
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
	EndpointName() *string
	SetEndpointName(val *string)
	EndpointNameInput() *string
	// Experimental.
	Fqn() *string
	GroupId() *string
	SetGroupId(val *string)
	GroupIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NetworkConnectivityConfigId() *string
	SetNetworkConnectivityConfigId(val *string)
	NetworkConnectivityConfigIdInput() *string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
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
	ResetConnectionState()
	ResetCreationTime()
	ResetDeactivated()
	ResetDeactivatedAt()
	ResetDomainNames()
	ResetEndpointName()
	ResetGroupId()
	ResetNetworkConnectivityConfigId()
	ResetResourceId()
	ResetRuleId()
	ResetUpdatedTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference
type jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ConnectionState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ConnectionStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) CreationTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) Deactivated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) DeactivatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deactivatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) DeactivatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deactivatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) DeactivatedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) DomainNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) DomainNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) EndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) GroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) NetworkConnectivityConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) NetworkConnectivityConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) RuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) RuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) UpdatedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) UpdatedTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedTimeInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksMwsNetworkConnectivityConfig.DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference_Override(d DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksMwsNetworkConnectivityConfig.DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetConnectionState(val *string) {
	if err := j.validateSetConnectionStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionState",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetCreationTime(val *float64) {
	if err := j.validateSetCreationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetDeactivated(val interface{}) {
	if err := j.validateSetDeactivatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deactivated",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetDeactivatedAt(val *float64) {
	if err := j.validateSetDeactivatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deactivatedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetDomainNames(val *[]*string) {
	if err := j.validateSetDomainNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainNames",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetEndpointName(val *string) {
	if err := j.validateSetEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetGroupId(val *string) {
	if err := j.validateSetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetNetworkConnectivityConfigId(val *string) {
	if err := j.validateSetNetworkConnectivityConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConnectivityConfigId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetRuleId(val *string) {
	if err := j.validateSetRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference)SetUpdatedTime(val *float64) {
	if err := j.validateSetUpdatedTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedTime",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResetConnectionState() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResetCreationTime() {
	_jsii_.InvokeVoid(
		d,
		"resetCreationTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResetDeactivated() {
	_jsii_.InvokeVoid(
		d,
		"resetDeactivated",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResetDeactivatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetDeactivatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResetDomainNames() {
	_jsii_.InvokeVoid(
		d,
		"resetDomainNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResetEndpointName() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResetGroupId() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResetNetworkConnectivityConfigId() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkConnectivityConfigId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResetRuleId() {
	_jsii_.InvokeVoid(
		d,
		"resetRuleId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ResetUpdatedTime() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksMwsNetworkConnectivityConfigEgressConfigTargetRulesAzurePrivateEndpointRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

