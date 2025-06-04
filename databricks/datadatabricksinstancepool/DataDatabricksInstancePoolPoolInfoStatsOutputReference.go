// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksinstancepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/datadatabricksinstancepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksInstancePoolPoolInfoStatsOutputReference interface {
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
	IdleCount() *float64
	SetIdleCount(val *float64)
	IdleCountInput() *float64
	InternalValue() *DataDatabricksInstancePoolPoolInfoStats
	SetInternalValue(val *DataDatabricksInstancePoolPoolInfoStats)
	PendingIdleCount() *float64
	SetPendingIdleCount(val *float64)
	PendingIdleCountInput() *float64
	PendingUsedCount() *float64
	SetPendingUsedCount(val *float64)
	PendingUsedCountInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsedCount() *float64
	SetUsedCount(val *float64)
	UsedCountInput() *float64
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
	ResetIdleCount()
	ResetPendingIdleCount()
	ResetPendingUsedCount()
	ResetUsedCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksInstancePoolPoolInfoStatsOutputReference
type jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) IdleCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) IdleCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) InternalValue() *DataDatabricksInstancePoolPoolInfoStats {
	var returns *DataDatabricksInstancePoolPoolInfoStats
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) PendingIdleCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pendingIdleCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) PendingIdleCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pendingIdleCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) PendingUsedCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pendingUsedCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) PendingUsedCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pendingUsedCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) UsedCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) UsedCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedCountInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksInstancePoolPoolInfoStatsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksInstancePoolPoolInfoStatsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksInstancePoolPoolInfoStatsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksInstancePool.DataDatabricksInstancePoolPoolInfoStatsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksInstancePoolPoolInfoStatsOutputReference_Override(d DataDatabricksInstancePoolPoolInfoStatsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksInstancePool.DataDatabricksInstancePoolPoolInfoStatsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference)SetIdleCount(val *float64) {
	if err := j.validateSetIdleCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleCount",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference)SetInternalValue(val *DataDatabricksInstancePoolPoolInfoStats) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference)SetPendingIdleCount(val *float64) {
	if err := j.validateSetPendingIdleCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pendingIdleCount",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference)SetPendingUsedCount(val *float64) {
	if err := j.validateSetPendingUsedCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pendingUsedCount",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference)SetUsedCount(val *float64) {
	if err := j.validateSetUsedCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usedCount",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) ResetIdleCount() {
	_jsii_.InvokeVoid(
		d,
		"resetIdleCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) ResetPendingIdleCount() {
	_jsii_.InvokeVoid(
		d,
		"resetPendingIdleCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) ResetPendingUsedCount() {
	_jsii_.InvokeVoid(
		d,
		"resetPendingUsedCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) ResetUsedCount() {
	_jsii_.InvokeVoid(
		d,
		"resetUsedCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoStatsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

