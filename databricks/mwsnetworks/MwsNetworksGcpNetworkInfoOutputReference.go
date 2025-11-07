// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/mwsnetworks/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsNetworksGcpNetworkInfoOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MwsNetworksGcpNetworkInfo
	SetInternalValue(val *MwsNetworksGcpNetworkInfo)
	NetworkProjectId() *string
	SetNetworkProjectId(val *string)
	NetworkProjectIdInput() *string
	PodIpRangeName() *string
	SetPodIpRangeName(val *string)
	PodIpRangeNameInput() *string
	ServiceIpRangeName() *string
	SetServiceIpRangeName(val *string)
	ServiceIpRangeNameInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	SubnetRegion() *string
	SetSubnetRegion(val *string)
	SubnetRegionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	ResetPodIpRangeName()
	ResetServiceIpRangeName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MwsNetworksGcpNetworkInfoOutputReference
type jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) InternalValue() *MwsNetworksGcpNetworkInfo {
	var returns *MwsNetworksGcpNetworkInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) NetworkProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) NetworkProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) PodIpRangeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podIpRangeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) PodIpRangeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podIpRangeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) ServiceIpRangeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpRangeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) ServiceIpRangeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpRangeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) SubnetRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) SubnetRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


func NewMwsNetworksGcpNetworkInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MwsNetworksGcpNetworkInfoOutputReference {
	_init_.Initialize()

	if err := validateNewMwsNetworksGcpNetworkInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsNetworks.MwsNetworksGcpNetworkInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMwsNetworksGcpNetworkInfoOutputReference_Override(m MwsNetworksGcpNetworkInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsNetworks.MwsNetworksGcpNetworkInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference)SetInternalValue(val *MwsNetworksGcpNetworkInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference)SetNetworkProjectId(val *string) {
	if err := j.validateSetNetworkProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkProjectId",
		val,
	)
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference)SetPodIpRangeName(val *string) {
	if err := j.validateSetPodIpRangeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podIpRangeName",
		val,
	)
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference)SetServiceIpRangeName(val *string) {
	if err := j.validateSetServiceIpRangeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceIpRangeName",
		val,
	)
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference)SetSubnetRegion(val *string) {
	if err := j.validateSetSubnetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetRegion",
		val,
	)
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) ResetPodIpRangeName() {
	_jsii_.InvokeVoid(
		m,
		"resetPodIpRangeName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) ResetServiceIpRangeName() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceIpRangeName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MwsNetworksGcpNetworkInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

