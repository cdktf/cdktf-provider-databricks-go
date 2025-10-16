// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabaseinstances

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksdatabaseinstances/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference interface {
	cdktf.ComplexObject
	Capacity() *string
	SetCapacity(val *string)
	CapacityInput() *string
	ChildInstanceRefs() DataDatabricksDatabaseInstancesDatabaseInstancesChildInstanceRefsList
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
	CreationTime() *string
	Creator() *string
	CustomTags() DataDatabricksDatabaseInstancesDatabaseInstancesCustomTagsList
	CustomTagsInput() interface{}
	EffectiveCapacity() *string
	EffectiveCustomTags() DataDatabricksDatabaseInstancesDatabaseInstancesEffectiveCustomTagsList
	EffectiveEnablePgNativeLogin() cdktf.IResolvable
	EffectiveEnableReadableSecondaries() cdktf.IResolvable
	EffectiveNodeCount() *float64
	EffectiveRetentionWindowInDays() *float64
	EffectiveStopped() cdktf.IResolvable
	EffectiveUsagePolicyId() *string
	EnablePgNativeLogin() interface{}
	SetEnablePgNativeLogin(val interface{})
	EnablePgNativeLoginInput() interface{}
	EnableReadableSecondaries() interface{}
	SetEnableReadableSecondaries(val interface{})
	EnableReadableSecondariesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksDatabaseInstancesDatabaseInstances
	SetInternalValue(val *DataDatabricksDatabaseInstancesDatabaseInstances)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	ParentInstanceRef() DataDatabricksDatabaseInstancesDatabaseInstancesParentInstanceRefOutputReference
	ParentInstanceRefInput() interface{}
	PgVersion() *string
	ReadOnlyDns() *string
	ReadWriteDns() *string
	RetentionWindowInDays() *float64
	SetRetentionWindowInDays(val *float64)
	RetentionWindowInDaysInput() *float64
	State() *string
	Stopped() interface{}
	SetStopped(val interface{})
	StoppedInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uid() *string
	UsagePolicyId() *string
	SetUsagePolicyId(val *string)
	UsagePolicyIdInput() *string
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
	PutCustomTags(value interface{})
	PutParentInstanceRef(value *DataDatabricksDatabaseInstancesDatabaseInstancesParentInstanceRef)
	ResetCapacity()
	ResetCustomTags()
	ResetEnablePgNativeLogin()
	ResetEnableReadableSecondaries()
	ResetNodeCount()
	ResetParentInstanceRef()
	ResetRetentionWindowInDays()
	ResetStopped()
	ResetUsagePolicyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference
type jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) Capacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) CapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ChildInstanceRefs() DataDatabricksDatabaseInstancesDatabaseInstancesChildInstanceRefsList {
	var returns DataDatabricksDatabaseInstancesDatabaseInstancesChildInstanceRefsList
	_jsii_.Get(
		j,
		"childInstanceRefs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) CustomTags() DataDatabricksDatabaseInstancesDatabaseInstancesCustomTagsList {
	var returns DataDatabricksDatabaseInstancesDatabaseInstancesCustomTagsList
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) CustomTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EffectiveCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EffectiveCustomTags() DataDatabricksDatabaseInstancesDatabaseInstancesEffectiveCustomTagsList {
	var returns DataDatabricksDatabaseInstancesDatabaseInstancesEffectiveCustomTagsList
	_jsii_.Get(
		j,
		"effectiveCustomTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EffectiveEnablePgNativeLogin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"effectiveEnablePgNativeLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EffectiveEnableReadableSecondaries() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"effectiveEnableReadableSecondaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EffectiveNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"effectiveNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EffectiveRetentionWindowInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"effectiveRetentionWindowInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EffectiveStopped() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"effectiveStopped",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EffectiveUsagePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveUsagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EnablePgNativeLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePgNativeLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EnablePgNativeLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePgNativeLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EnableReadableSecondaries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableReadableSecondaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) EnableReadableSecondariesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableReadableSecondariesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) InternalValue() *DataDatabricksDatabaseInstancesDatabaseInstances {
	var returns *DataDatabricksDatabaseInstancesDatabaseInstances
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ParentInstanceRef() DataDatabricksDatabaseInstancesDatabaseInstancesParentInstanceRefOutputReference {
	var returns DataDatabricksDatabaseInstancesDatabaseInstancesParentInstanceRefOutputReference
	_jsii_.Get(
		j,
		"parentInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ParentInstanceRefInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parentInstanceRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) PgVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ReadOnlyDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readOnlyDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ReadWriteDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readWriteDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) RetentionWindowInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionWindowInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) RetentionWindowInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionWindowInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) Stopped() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stopped",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) StoppedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) UsagePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) UsagePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyIdInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksDatabaseInstancesDatabaseInstancesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksDatabaseInstancesDatabaseInstancesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDatabaseInstances.DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksDatabaseInstancesDatabaseInstancesOutputReference_Override(d DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDatabaseInstances.DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetCapacity(val *string) {
	if err := j.validateSetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetEnablePgNativeLogin(val interface{}) {
	if err := j.validateSetEnablePgNativeLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePgNativeLogin",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetEnableReadableSecondaries(val interface{}) {
	if err := j.validateSetEnableReadableSecondariesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableReadableSecondaries",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetInternalValue(val *DataDatabricksDatabaseInstancesDatabaseInstances) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetRetentionWindowInDays(val *float64) {
	if err := j.validateSetRetentionWindowInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionWindowInDays",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetStopped(val interface{}) {
	if err := j.validateSetStoppedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stopped",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference)SetUsagePolicyId(val *string) {
	if err := j.validateSetUsagePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usagePolicyId",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) PutCustomTags(value interface{}) {
	if err := d.validatePutCustomTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) PutParentInstanceRef(value *DataDatabricksDatabaseInstancesDatabaseInstancesParentInstanceRef) {
	if err := d.validatePutParentInstanceRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putParentInstanceRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ResetCapacity() {
	_jsii_.InvokeVoid(
		d,
		"resetCapacity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ResetEnablePgNativeLogin() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePgNativeLogin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ResetEnableReadableSecondaries() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableReadableSecondaries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ResetParentInstanceRef() {
	_jsii_.InvokeVoid(
		d,
		"resetParentInstanceRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ResetRetentionWindowInDays() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionWindowInDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ResetStopped() {
	_jsii_.InvokeVoid(
		d,
		"resetStopped",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ResetUsagePolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetUsagePolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksDatabaseInstancesDatabaseInstancesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

