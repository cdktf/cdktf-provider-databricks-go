// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomscleanroom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscleanroomscleanroom/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference interface {
	cdktf.ComplexObject
	CentralCleanRoomId() *string
	CloudVendor() *string
	SetCloudVendor(val *string)
	CloudVendorInput() *string
	Collaborators() DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsList
	CollaboratorsInput() interface{}
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
	ComplianceSecurityProfile() DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoComplianceSecurityProfileOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Creator() DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference
	EgressNetworkPolicy() DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyOutputReference
	EgressNetworkPolicyInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
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
	PutCollaborators(value interface{})
	PutEgressNetworkPolicy(value *DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicy)
	ResetCloudVendor()
	ResetCollaborators()
	ResetEgressNetworkPolicy()
	ResetRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference
type jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) CentralCleanRoomId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centralCleanRoomId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) CloudVendor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudVendor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) CloudVendorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudVendorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) Collaborators() DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsList {
	var returns DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsList
	_jsii_.Get(
		j,
		"collaborators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) CollaboratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collaboratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ComplianceSecurityProfile() DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoComplianceSecurityProfileOutputReference {
	var returns DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoComplianceSecurityProfileOutputReference
	_jsii_.Get(
		j,
		"complianceSecurityProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) Creator() DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference {
	var returns DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) EgressNetworkPolicy() DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyOutputReference {
	var returns DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyOutputReference
	_jsii_.Get(
		j,
		"egressNetworkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) EgressNetworkPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressNetworkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomsCleanRoom.DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference_Override(d DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomsCleanRoom.DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetCloudVendor(val *string) {
	if err := j.validateSetCloudVendorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudVendor",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) PutCollaborators(value interface{}) {
	if err := d.validatePutCollaboratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCollaborators",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) PutEgressNetworkPolicy(value *DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicy) {
	if err := d.validatePutEgressNetworkPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEgressNetworkPolicy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ResetCloudVendor() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudVendor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ResetCollaborators() {
	_jsii_.InvokeVoid(
		d,
		"resetCollaborators",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ResetEgressNetworkPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetEgressNetworkPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

