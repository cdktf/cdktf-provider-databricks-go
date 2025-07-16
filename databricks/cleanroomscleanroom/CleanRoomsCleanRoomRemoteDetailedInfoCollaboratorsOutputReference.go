// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomscleanroom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/cleanroomscleanroom/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference interface {
	cdktf.ComplexObject
	CollaboratorAlias() *string
	SetCollaboratorAlias(val *string)
	CollaboratorAliasInput() *string
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
	DisplayName() *string
	// Experimental.
	Fqn() *string
	GlobalMetastoreId() *string
	SetGlobalMetastoreId(val *string)
	GlobalMetastoreIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	InviteRecipientEmail() *string
	SetInviteRecipientEmail(val *string)
	InviteRecipientEmailInput() *string
	InviteRecipientWorkspaceId() *float64
	SetInviteRecipientWorkspaceId(val *float64)
	InviteRecipientWorkspaceIdInput() *float64
	OrganizationName() *string
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
	ResetGlobalMetastoreId()
	ResetInviteRecipientEmail()
	ResetInviteRecipientWorkspaceId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference
type jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) CollaboratorAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) CollaboratorAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) GlobalMetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalMetastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) GlobalMetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalMetastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) InviteRecipientEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inviteRecipientEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) InviteRecipientEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inviteRecipientEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) InviteRecipientWorkspaceId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inviteRecipientWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) InviteRecipientWorkspaceIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inviteRecipientWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference {
	_init_.Initialize()

	if err := validateNewCleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomsCleanRoom.CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference_Override(c CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomsCleanRoom.CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference)SetCollaboratorAlias(val *string) {
	if err := j.validateSetCollaboratorAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collaboratorAlias",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference)SetGlobalMetastoreId(val *string) {
	if err := j.validateSetGlobalMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalMetastoreId",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference)SetInviteRecipientEmail(val *string) {
	if err := j.validateSetInviteRecipientEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inviteRecipientEmail",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference)SetInviteRecipientWorkspaceId(val *float64) {
	if err := j.validateSetInviteRecipientWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inviteRecipientWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) ResetGlobalMetastoreId() {
	_jsii_.InvokeVoid(
		c,
		"resetGlobalMetastoreId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) ResetInviteRecipientEmail() {
	_jsii_.InvokeVoid(
		c,
		"resetInviteRecipientEmail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) ResetInviteRecipientWorkspaceId() {
	_jsii_.InvokeVoid(
		c,
		"resetInviteRecipientWorkspaceId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoCollaboratorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

