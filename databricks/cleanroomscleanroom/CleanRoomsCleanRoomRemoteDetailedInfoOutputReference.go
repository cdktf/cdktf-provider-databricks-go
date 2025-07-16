// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomscleanroom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/cleanroomscleanroom/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanRoomsCleanRoomRemoteDetailedInfoOutputReference interface {
	cdktf.ComplexObject
	CentralCleanRoomId() *string
	CloudVendor() *string
	SetCloudVendor(val *string)
	CloudVendorInput() *string
	Collaborators() CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsList
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
	ComplianceSecurityProfile() CleanRoomsCleanRoomRemoteDetailedInfoComplianceSecurityProfileOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Creator() CleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference
	EgressNetworkPolicy() CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyOutputReference
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
	PutEgressNetworkPolicy(value *CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicy)
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

// The jsii proxy struct for CleanRoomsCleanRoomRemoteDetailedInfoOutputReference
type jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) CentralCleanRoomId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"centralCleanRoomId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) CloudVendor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudVendor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) CloudVendorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudVendorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) Collaborators() CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsList {
	var returns CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsList
	_jsii_.Get(
		j,
		"collaborators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) CollaboratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collaboratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ComplianceSecurityProfile() CleanRoomsCleanRoomRemoteDetailedInfoComplianceSecurityProfileOutputReference {
	var returns CleanRoomsCleanRoomRemoteDetailedInfoComplianceSecurityProfileOutputReference
	_jsii_.Get(
		j,
		"complianceSecurityProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) Creator() CleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference {
	var returns CleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) EgressNetworkPolicy() CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyOutputReference {
	var returns CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyOutputReference
	_jsii_.Get(
		j,
		"egressNetworkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) EgressNetworkPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressNetworkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanRoomsCleanRoomRemoteDetailedInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CleanRoomsCleanRoomRemoteDetailedInfoOutputReference {
	_init_.Initialize()

	if err := validateNewCleanRoomsCleanRoomRemoteDetailedInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomsCleanRoom.CleanRoomsCleanRoomRemoteDetailedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanRoomsCleanRoomRemoteDetailedInfoOutputReference_Override(c CleanRoomsCleanRoomRemoteDetailedInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomsCleanRoom.CleanRoomsCleanRoomRemoteDetailedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetCloudVendor(val *string) {
	if err := j.validateSetCloudVendorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudVendor",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) PutCollaborators(value interface{}) {
	if err := c.validatePutCollaboratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCollaborators",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) PutEgressNetworkPolicy(value *CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicy) {
	if err := c.validatePutEgressNetworkPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEgressNetworkPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ResetCloudVendor() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudVendor",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ResetCollaborators() {
	_jsii_.InvokeVoid(
		c,
		"resetCollaborators",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ResetEgressNetworkPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetEgressNetworkPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

