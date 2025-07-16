// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomscleanroom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscleanroomscleanroom/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference interface {
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
	InternalValue() *DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreator
	SetInternalValue(val *DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreator)
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

// The jsii proxy struct for DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference
type jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) CollaboratorAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) CollaboratorAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) GlobalMetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalMetastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) GlobalMetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalMetastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) InternalValue() *DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreator {
	var returns *DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreator
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) InviteRecipientEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inviteRecipientEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) InviteRecipientEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inviteRecipientEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) InviteRecipientWorkspaceId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inviteRecipientWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) InviteRecipientWorkspaceIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inviteRecipientWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomsCleanRoom.DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference_Override(d DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCleanRoomsCleanRoom.DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference)SetCollaboratorAlias(val *string) {
	if err := j.validateSetCollaboratorAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collaboratorAlias",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference)SetGlobalMetastoreId(val *string) {
	if err := j.validateSetGlobalMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalMetastoreId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference)SetInternalValue(val *DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreator) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference)SetInviteRecipientEmail(val *string) {
	if err := j.validateSetInviteRecipientEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inviteRecipientEmail",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference)SetInviteRecipientWorkspaceId(val *float64) {
	if err := j.validateSetInviteRecipientWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inviteRecipientWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) ResetGlobalMetastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetGlobalMetastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) ResetInviteRecipientEmail() {
	_jsii_.InvokeVoid(
		d,
		"resetInviteRecipientEmail",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) ResetInviteRecipientWorkspaceId() {
	_jsii_.InvokeVoid(
		d,
		"resetInviteRecipientWorkspaceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoCreatorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

