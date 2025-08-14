// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/notificationdestination/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationDestinationConfigSlackOutputReference interface {
	cdktf.ComplexObject
	ChannelId() *string
	SetChannelId(val *string)
	ChannelIdInput() *string
	ChannelIdSet() interface{}
	SetChannelIdSet(val interface{})
	ChannelIdSetInput() interface{}
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
	InternalValue() *NotificationDestinationConfigSlack
	SetInternalValue(val *NotificationDestinationConfigSlack)
	OauthToken() *string
	SetOauthToken(val *string)
	OauthTokenInput() *string
	OauthTokenSet() interface{}
	SetOauthTokenSet(val interface{})
	OauthTokenSetInput() interface{}
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetChannelId()
	ResetChannelIdSet()
	ResetOauthToken()
	ResetOauthTokenSet()
	ResetUrl()
	ResetUrlSet()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotificationDestinationConfigSlackOutputReference
type jsiiProxy_NotificationDestinationConfigSlackOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ChannelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ChannelIdSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelIdSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ChannelIdSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelIdSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) InternalValue() *NotificationDestinationConfigSlack {
	var returns *NotificationDestinationConfigSlack
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) OauthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) OauthTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) OauthTokenSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauthTokenSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) OauthTokenSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauthTokenSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) UrlSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference) UrlSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlSetInput",
		&returns,
	)
	return returns
}


func NewNotificationDestinationConfigSlackOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotificationDestinationConfigSlackOutputReference {
	_init_.Initialize()

	if err := validateNewNotificationDestinationConfigSlackOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotificationDestinationConfigSlackOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.notificationDestination.NotificationDestinationConfigSlackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotificationDestinationConfigSlackOutputReference_Override(n NotificationDestinationConfigSlackOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.notificationDestination.NotificationDestinationConfigSlackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference)SetChannelId(val *string) {
	if err := j.validateSetChannelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelId",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference)SetChannelIdSet(val interface{}) {
	if err := j.validateSetChannelIdSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelIdSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference)SetInternalValue(val *NotificationDestinationConfigSlack) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference)SetOauthToken(val *string) {
	if err := j.validateSetOauthTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthToken",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference)SetOauthTokenSet(val interface{}) {
	if err := j.validateSetOauthTokenSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthTokenSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigSlackOutputReference)SetUrlSet(val interface{}) {
	if err := j.validateSetUrlSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlSet",
		val,
	)
}

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ResetChannelId() {
	_jsii_.InvokeVoid(
		n,
		"resetChannelId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ResetChannelIdSet() {
	_jsii_.InvokeVoid(
		n,
		"resetChannelIdSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ResetOauthToken() {
	_jsii_.InvokeVoid(
		n,
		"resetOauthToken",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ResetOauthTokenSet() {
	_jsii_.InvokeVoid(
		n,
		"resetOauthTokenSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		n,
		"resetUrl",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ResetUrlSet() {
	_jsii_.InvokeVoid(
		n,
		"resetUrlSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigSlackOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

