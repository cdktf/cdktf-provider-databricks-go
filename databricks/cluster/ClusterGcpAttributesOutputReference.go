// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/cluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterGcpAttributesOutputReference interface {
	cdktf.ComplexObject
	Availability() *string
	SetAvailability(val *string)
	AvailabilityInput() *string
	BootDiskSize() *float64
	SetBootDiskSize(val *float64)
	BootDiskSizeInput() *float64
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
	GoogleServiceAccount() *string
	SetGoogleServiceAccount(val *string)
	GoogleServiceAccountInput() *string
	InternalValue() *ClusterGcpAttributes
	SetInternalValue(val *ClusterGcpAttributes)
	LocalSsdCount() *float64
	SetLocalSsdCount(val *float64)
	LocalSsdCountInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsePreemptibleExecutors() interface{}
	SetUsePreemptibleExecutors(val interface{})
	UsePreemptibleExecutorsInput() interface{}
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	ResetAvailability()
	ResetBootDiskSize()
	ResetGoogleServiceAccount()
	ResetLocalSsdCount()
	ResetUsePreemptibleExecutors()
	ResetZoneId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClusterGcpAttributesOutputReference
type jsiiProxy_ClusterGcpAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) Availability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) AvailabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) BootDiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) BootDiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) GoogleServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) GoogleServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) InternalValue() *ClusterGcpAttributes {
	var returns *ClusterGcpAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) LocalSsdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) LocalSsdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) UsePreemptibleExecutors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePreemptibleExecutors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) UsePreemptibleExecutorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePreemptibleExecutorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


func NewClusterGcpAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ClusterGcpAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewClusterGcpAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClusterGcpAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.cluster.ClusterGcpAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewClusterGcpAttributesOutputReference_Override(c ClusterGcpAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.cluster.ClusterGcpAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference)SetAvailability(val *string) {
	if err := j.validateSetAvailabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availability",
		val,
	)
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference)SetBootDiskSize(val *float64) {
	if err := j.validateSetBootDiskSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSize",
		val,
	)
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference)SetGoogleServiceAccount(val *string) {
	if err := j.validateSetGoogleServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleServiceAccount",
		val,
	)
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference)SetInternalValue(val *ClusterGcpAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference)SetLocalSsdCount(val *float64) {
	if err := j.validateSetLocalSsdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdCount",
		val,
	)
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference)SetUsePreemptibleExecutors(val interface{}) {
	if err := j.validateSetUsePreemptibleExecutorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePreemptibleExecutors",
		val,
	)
}

func (j *jsiiProxy_ClusterGcpAttributesOutputReference)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) ResetAvailability() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailability",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) ResetBootDiskSize() {
	_jsii_.InvokeVoid(
		c,
		"resetBootDiskSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) ResetGoogleServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetGoogleServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) ResetLocalSsdCount() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalSsdCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) ResetUsePreemptibleExecutors() {
	_jsii_.InvokeVoid(
		c,
		"resetUsePreemptibleExecutors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) ResetZoneId() {
	_jsii_.InvokeVoid(
		c,
		"resetZoneId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClusterGcpAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

