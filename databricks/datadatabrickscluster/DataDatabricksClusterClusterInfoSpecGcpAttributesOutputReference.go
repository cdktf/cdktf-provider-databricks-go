// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference interface {
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
	FirstOnDemand() *float64
	SetFirstOnDemand(val *float64)
	FirstOnDemandInput() *float64
	// Experimental.
	Fqn() *string
	GoogleServiceAccount() *string
	SetGoogleServiceAccount(val *string)
	GoogleServiceAccountInput() *string
	InternalValue() *DataDatabricksClusterClusterInfoSpecGcpAttributes
	SetInternalValue(val *DataDatabricksClusterClusterInfoSpecGcpAttributes)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAvailability()
	ResetBootDiskSize()
	ResetFirstOnDemand()
	ResetGoogleServiceAccount()
	ResetLocalSsdCount()
	ResetUsePreemptibleExecutors()
	ResetZoneId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference
type jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) Availability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) AvailabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) BootDiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) BootDiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) FirstOnDemand() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstOnDemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) FirstOnDemandInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstOnDemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) GoogleServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) GoogleServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) InternalValue() *DataDatabricksClusterClusterInfoSpecGcpAttributes {
	var returns *DataDatabricksClusterClusterInfoSpecGcpAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) LocalSsdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) LocalSsdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) UsePreemptibleExecutors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePreemptibleExecutors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) UsePreemptibleExecutorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePreemptibleExecutorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksClusterClusterInfoSpecGcpAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCluster.DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference_Override(d DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCluster.DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetAvailability(val *string) {
	if err := j.validateSetAvailabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availability",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetBootDiskSize(val *float64) {
	if err := j.validateSetBootDiskSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSize",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetFirstOnDemand(val *float64) {
	if err := j.validateSetFirstOnDemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstOnDemand",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetGoogleServiceAccount(val *string) {
	if err := j.validateSetGoogleServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleServiceAccount",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetInternalValue(val *DataDatabricksClusterClusterInfoSpecGcpAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetLocalSsdCount(val *float64) {
	if err := j.validateSetLocalSsdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdCount",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetUsePreemptibleExecutors(val interface{}) {
	if err := j.validateSetUsePreemptibleExecutorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePreemptibleExecutors",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ResetAvailability() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailability",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ResetBootDiskSize() {
	_jsii_.InvokeVoid(
		d,
		"resetBootDiskSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ResetFirstOnDemand() {
	_jsii_.InvokeVoid(
		d,
		"resetFirstOnDemand",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ResetGoogleServiceAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetGoogleServiceAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ResetLocalSsdCount() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalSsdCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ResetUsePreemptibleExecutors() {
	_jsii_.InvokeVoid(
		d,
		"resetUsePreemptibleExecutors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ResetZoneId() {
	_jsii_.InvokeVoid(
		d,
		"resetZoneId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

