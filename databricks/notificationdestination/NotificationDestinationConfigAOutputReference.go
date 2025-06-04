// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/notificationdestination/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationDestinationConfigAOutputReference interface {
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
	Email() NotificationDestinationConfigEmailOutputReference
	EmailInput() *NotificationDestinationConfigEmail
	// Experimental.
	Fqn() *string
	GenericWebhook() NotificationDestinationConfigGenericWebhookOutputReference
	GenericWebhookInput() *NotificationDestinationConfigGenericWebhook
	InternalValue() *NotificationDestinationConfigA
	SetInternalValue(val *NotificationDestinationConfigA)
	MicrosoftTeams() NotificationDestinationConfigMicrosoftTeamsOutputReference
	MicrosoftTeamsInput() *NotificationDestinationConfigMicrosoftTeams
	Pagerduty() NotificationDestinationConfigPagerdutyOutputReference
	PagerdutyInput() *NotificationDestinationConfigPagerduty
	Slack() NotificationDestinationConfigSlackOutputReference
	SlackInput() *NotificationDestinationConfigSlack
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
	PutEmail(value *NotificationDestinationConfigEmail)
	PutGenericWebhook(value *NotificationDestinationConfigGenericWebhook)
	PutMicrosoftTeams(value *NotificationDestinationConfigMicrosoftTeams)
	PutPagerduty(value *NotificationDestinationConfigPagerduty)
	PutSlack(value *NotificationDestinationConfigSlack)
	ResetEmail()
	ResetGenericWebhook()
	ResetMicrosoftTeams()
	ResetPagerduty()
	ResetSlack()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotificationDestinationConfigAOutputReference
type jsiiProxy_NotificationDestinationConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) Email() NotificationDestinationConfigEmailOutputReference {
	var returns NotificationDestinationConfigEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) EmailInput() *NotificationDestinationConfigEmail {
	var returns *NotificationDestinationConfigEmail
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) GenericWebhook() NotificationDestinationConfigGenericWebhookOutputReference {
	var returns NotificationDestinationConfigGenericWebhookOutputReference
	_jsii_.Get(
		j,
		"genericWebhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) GenericWebhookInput() *NotificationDestinationConfigGenericWebhook {
	var returns *NotificationDestinationConfigGenericWebhook
	_jsii_.Get(
		j,
		"genericWebhookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) InternalValue() *NotificationDestinationConfigA {
	var returns *NotificationDestinationConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) MicrosoftTeams() NotificationDestinationConfigMicrosoftTeamsOutputReference {
	var returns NotificationDestinationConfigMicrosoftTeamsOutputReference
	_jsii_.Get(
		j,
		"microsoftTeams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) MicrosoftTeamsInput() *NotificationDestinationConfigMicrosoftTeams {
	var returns *NotificationDestinationConfigMicrosoftTeams
	_jsii_.Get(
		j,
		"microsoftTeamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) Pagerduty() NotificationDestinationConfigPagerdutyOutputReference {
	var returns NotificationDestinationConfigPagerdutyOutputReference
	_jsii_.Get(
		j,
		"pagerduty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) PagerdutyInput() *NotificationDestinationConfigPagerduty {
	var returns *NotificationDestinationConfigPagerduty
	_jsii_.Get(
		j,
		"pagerdutyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) Slack() NotificationDestinationConfigSlackOutputReference {
	var returns NotificationDestinationConfigSlackOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) SlackInput() *NotificationDestinationConfigSlack {
	var returns *NotificationDestinationConfigSlack
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNotificationDestinationConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotificationDestinationConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewNotificationDestinationConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotificationDestinationConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.notificationDestination.NotificationDestinationConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotificationDestinationConfigAOutputReference_Override(n NotificationDestinationConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.notificationDestination.NotificationDestinationConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference)SetInternalValue(val *NotificationDestinationConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) PutEmail(value *NotificationDestinationConfigEmail) {
	if err := n.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putEmail",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) PutGenericWebhook(value *NotificationDestinationConfigGenericWebhook) {
	if err := n.validatePutGenericWebhookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putGenericWebhook",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) PutMicrosoftTeams(value *NotificationDestinationConfigMicrosoftTeams) {
	if err := n.validatePutMicrosoftTeamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putMicrosoftTeams",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) PutPagerduty(value *NotificationDestinationConfigPagerduty) {
	if err := n.validatePutPagerdutyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putPagerduty",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) PutSlack(value *NotificationDestinationConfigSlack) {
	if err := n.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSlack",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		n,
		"resetEmail",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) ResetGenericWebhook() {
	_jsii_.InvokeVoid(
		n,
		"resetGenericWebhook",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) ResetMicrosoftTeams() {
	_jsii_.InvokeVoid(
		n,
		"resetMicrosoftTeams",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) ResetPagerduty() {
	_jsii_.InvokeVoid(
		n,
		"resetPagerduty",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		n,
		"resetSlack",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NotificationDestinationConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

