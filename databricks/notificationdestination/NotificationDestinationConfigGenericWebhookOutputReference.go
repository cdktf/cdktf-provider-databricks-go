// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/notificationdestination/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationDestinationConfigGenericWebhookOutputReference interface {
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
	InternalValue() *NotificationDestinationConfigGenericWebhook
	SetInternalValue(val *NotificationDestinationConfigGenericWebhook)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PasswordSet() interface{}
	SetPasswordSet(val interface{})
	PasswordSetInput() interface{}
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
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	UsernameSet() interface{}
	SetUsernameSet(val interface{})
	UsernameSetInput() interface{}
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
	ResetPassword()
	ResetPasswordSet()
	ResetUrl()
	ResetUrlSet()
	ResetUsername()
	ResetUsernameSet()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotificationDestinationConfigGenericWebhookOutputReference
type jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) InternalValue() *NotificationDestinationConfigGenericWebhook {
	var returns *NotificationDestinationConfigGenericWebhook
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) PasswordSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) PasswordSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) UrlSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) UrlSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) UsernameSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usernameSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) UsernameSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usernameSetInput",
		&returns,
	)
	return returns
}


func NewNotificationDestinationConfigGenericWebhookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotificationDestinationConfigGenericWebhookOutputReference {
	_init_.Initialize()

	if err := validateNewNotificationDestinationConfigGenericWebhookOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.notificationDestination.NotificationDestinationConfigGenericWebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotificationDestinationConfigGenericWebhookOutputReference_Override(n NotificationDestinationConfigGenericWebhookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.notificationDestination.NotificationDestinationConfigGenericWebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference)SetInternalValue(val *NotificationDestinationConfigGenericWebhook) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference)SetPasswordSet(val interface{}) {
	if err := j.validateSetPasswordSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference)SetUrlSet(val interface{}) {
	if err := j.validateSetUrlSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference)SetUsernameSet(val interface{}) {
	if err := j.validateSetUsernameSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameSet",
		val,
	)
}

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		n,
		"resetPassword",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) ResetPasswordSet() {
	_jsii_.InvokeVoid(
		n,
		"resetPasswordSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		n,
		"resetUrl",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) ResetUrlSet() {
	_jsii_.InvokeVoid(
		n,
		"resetUrlSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		n,
		"resetUsername",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) ResetUsernameSet() {
	_jsii_.InvokeVoid(
		n,
		"resetUsernameSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NotificationDestinationConfigGenericWebhookOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

