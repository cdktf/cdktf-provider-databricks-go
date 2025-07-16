// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomscleanroom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/cleanroomscleanroom/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference interface {
	cdktf.ComplexObject
	AllowedInternetDestinations() CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedInternetDestinationsList
	AllowedInternetDestinationsInput() interface{}
	AllowedStorageDestinations() CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinationsList
	AllowedStorageDestinationsInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogOnlyMode() CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyModeOutputReference
	LogOnlyModeInput() interface{}
	RestrictionMode() *string
	SetRestrictionMode(val *string)
	RestrictionModeInput() *string
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
	PutAllowedInternetDestinations(value interface{})
	PutAllowedStorageDestinations(value interface{})
	PutLogOnlyMode(value *CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyMode)
	ResetAllowedInternetDestinations()
	ResetAllowedStorageDestinations()
	ResetLogOnlyMode()
	ResetRestrictionMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference
type jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) AllowedInternetDestinations() CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedInternetDestinationsList {
	var returns CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedInternetDestinationsList
	_jsii_.Get(
		j,
		"allowedInternetDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) AllowedInternetDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedInternetDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) AllowedStorageDestinations() CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinationsList {
	var returns CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinationsList
	_jsii_.Get(
		j,
		"allowedStorageDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) AllowedStorageDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedStorageDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) LogOnlyMode() CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyModeOutputReference {
	var returns CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyModeOutputReference
	_jsii_.Get(
		j,
		"logOnlyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) LogOnlyModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logOnlyModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) RestrictionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) RestrictionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference {
	_init_.Initialize()

	if err := validateNewCleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomsCleanRoom.CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference_Override(c CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.cleanRoomsCleanRoom.CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetRestrictionMode(val *string) {
	if err := j.validateSetRestrictionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictionMode",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) PutAllowedInternetDestinations(value interface{}) {
	if err := c.validatePutAllowedInternetDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAllowedInternetDestinations",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) PutAllowedStorageDestinations(value interface{}) {
	if err := c.validatePutAllowedStorageDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAllowedStorageDestinations",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) PutLogOnlyMode(value *CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessLogOnlyMode) {
	if err := c.validatePutLogOnlyModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLogOnlyMode",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ResetAllowedInternetDestinations() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedInternetDestinations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ResetAllowedStorageDestinations() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedStorageDestinations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ResetLogOnlyMode() {
	_jsii_.InvokeVoid(
		c,
		"resetLogOnlyMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ResetRestrictionMode() {
	_jsii_.InvokeVoid(
		c,
		"resetRestrictionMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

