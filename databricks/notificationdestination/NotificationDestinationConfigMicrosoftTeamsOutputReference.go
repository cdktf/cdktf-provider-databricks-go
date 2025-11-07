// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/notificationdestination/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationDestinationConfigMicrosoftTeamsOutputReference interface {
	cdktf.ComplexObject
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	AppIdSet() interface{}
	SetAppIdSet(val interface{})
	AppIdSetInput() interface{}
	AuthSecret() *string
	SetAuthSecret(val *string)
	AuthSecretInput() *string
	AuthSecretSet() interface{}
	SetAuthSecretSet(val interface{})
	AuthSecretSetInput() interface{}
	ChannelUrl() *string
	SetChannelUrl(val *string)
	ChannelUrlInput() *string
	ChannelUrlSet() interface{}
	SetChannelUrlSet(val interface{})
	ChannelUrlSetInput() interface{}
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
	InternalValue() *NotificationDestinationConfigMicrosoftTeams
	SetInternalValue(val *NotificationDestinationConfigMicrosoftTeams)
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	TenantIdSet() interface{}
	SetTenantIdSet(val interface{})
	TenantIdSetInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	UrlSet() interface{}
	SetUrlSet(val interface{})
	UrlSetInput() interface{}
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
	ResetAppId()
	ResetAppIdSet()
	ResetAuthSecret()
	ResetAuthSecretSet()
	ResetChannelUrl()
	ResetChannelUrlSet()
	ResetTenantId()
	ResetTenantIdSet()
	ResetUrl()
	ResetUrlSet()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotificationDestinationConfigMicrosoftTeamsOutputReference
type jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) AppIdSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appIdSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) AppIdSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appIdSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) AuthSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) AuthSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) AuthSecretSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authSecretSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) AuthSecretSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authSecretSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ChannelUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ChannelUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ChannelUrlSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelUrlSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ChannelUrlSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelUrlSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) InternalValue() *NotificationDestinationConfigMicrosoftTeams {
	var returns *NotificationDestinationConfigMicrosoftTeams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) TenantIdSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tenantIdSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) TenantIdSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tenantIdSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) UrlSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) UrlSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlSetInput",
		&returns,
	)
	return returns
}


func NewNotificationDestinationConfigMicrosoftTeamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotificationDestinationConfigMicrosoftTeamsOutputReference {
	_init_.Initialize()

	if err := validateNewNotificationDestinationConfigMicrosoftTeamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.notificationDestination.NotificationDestinationConfigMicrosoftTeamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotificationDestinationConfigMicrosoftTeamsOutputReference_Override(n NotificationDestinationConfigMicrosoftTeamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.notificationDestination.NotificationDestinationConfigMicrosoftTeamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetAppId(val *string) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetAppIdSet(val interface{}) {
	if err := j.validateSetAppIdSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appIdSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetAuthSecret(val *string) {
	if err := j.validateSetAuthSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authSecret",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetAuthSecretSet(val interface{}) {
	if err := j.validateSetAuthSecretSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authSecretSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetChannelUrl(val *string) {
	if err := j.validateSetChannelUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelUrl",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetChannelUrlSet(val interface{}) {
	if err := j.validateSetChannelUrlSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelUrlSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetInternalValue(val *NotificationDestinationConfigMicrosoftTeams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetTenantIdSet(val interface{}) {
	if err := j.validateSetTenantIdSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantIdSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference)SetUrlSet(val interface{}) {
	if err := j.validateSetUrlSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlSet",
		val,
	)
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ResetAppId() {
	_jsii_.InvokeVoid(
		n,
		"resetAppId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ResetAppIdSet() {
	_jsii_.InvokeVoid(
		n,
		"resetAppIdSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ResetAuthSecret() {
	_jsii_.InvokeVoid(
		n,
		"resetAuthSecret",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ResetAuthSecretSet() {
	_jsii_.InvokeVoid(
		n,
		"resetAuthSecretSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ResetChannelUrl() {
	_jsii_.InvokeVoid(
		n,
		"resetChannelUrl",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ResetChannelUrlSet() {
	_jsii_.InvokeVoid(
		n,
		"resetChannelUrlSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		n,
		"resetTenantId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ResetTenantIdSet() {
	_jsii_.InvokeVoid(
		n,
		"resetTenantIdSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		n,
		"resetUrl",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ResetUrlSet() {
	_jsii_.InvokeVoid(
		n,
		"resetUrlSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigMicrosoftTeamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

