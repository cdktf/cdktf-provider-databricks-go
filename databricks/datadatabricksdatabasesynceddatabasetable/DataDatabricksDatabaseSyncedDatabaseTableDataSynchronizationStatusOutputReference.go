// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksdatabasesynceddatabasetable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference interface {
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
	ContinuousUpdateStatus() DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusOutputReference
	ContinuousUpdateStatusInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DetailedState() *string
	FailedStatus() DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatusOutputReference
	FailedStatusInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatus
	SetInternalValue(val *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatus)
	LastSync() DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncOutputReference
	Message() *string
	PipelineId() *string
	ProvisioningStatus() DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusOutputReference
	ProvisioningStatusInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TriggeredUpdateStatus() DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusOutputReference
	TriggeredUpdateStatusInput() interface{}
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
	PutContinuousUpdateStatus(value *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatus)
	PutFailedStatus(value *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatus)
	PutProvisioningStatus(value *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatus)
	PutTriggeredUpdateStatus(value *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatus)
	ResetContinuousUpdateStatus()
	ResetFailedStatus()
	ResetProvisioningStatus()
	ResetTriggeredUpdateStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference
type jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ContinuousUpdateStatus() DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusOutputReference {
	var returns DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusOutputReference
	_jsii_.Get(
		j,
		"continuousUpdateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ContinuousUpdateStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousUpdateStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) DetailedState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detailedState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) FailedStatus() DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatusOutputReference {
	var returns DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatusOutputReference
	_jsii_.Get(
		j,
		"failedStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) FailedStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) InternalValue() *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatus {
	var returns *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatus
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) LastSync() DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncOutputReference {
	var returns DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncOutputReference
	_jsii_.Get(
		j,
		"lastSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) PipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ProvisioningStatus() DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusOutputReference {
	var returns DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusOutputReference
	_jsii_.Get(
		j,
		"provisioningStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ProvisioningStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisioningStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) TriggeredUpdateStatus() DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusOutputReference {
	var returns DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusOutputReference
	_jsii_.Get(
		j,
		"triggeredUpdateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) TriggeredUpdateStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggeredUpdateStatusInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference_Override(d DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference)SetInternalValue(val *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatus) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) PutContinuousUpdateStatus(value *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatus) {
	if err := d.validatePutContinuousUpdateStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContinuousUpdateStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) PutFailedStatus(value *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatus) {
	if err := d.validatePutFailedStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFailedStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) PutProvisioningStatus(value *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatus) {
	if err := d.validatePutProvisioningStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvisioningStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) PutTriggeredUpdateStatus(value *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatus) {
	if err := d.validatePutTriggeredUpdateStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTriggeredUpdateStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ResetContinuousUpdateStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetContinuousUpdateStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ResetFailedStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetFailedStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ResetProvisioningStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetProvisioningStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ResetTriggeredUpdateStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetTriggeredUpdateStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

