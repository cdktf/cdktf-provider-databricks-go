// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasesynceddatabasetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/databasesynceddatabasetable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference interface {
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
	ContinuousUpdateStatus() DatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusOutputReference
	ContinuousUpdateStatusInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DetailedState() *string
	FailedStatus() DatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatusOutputReference
	FailedStatusInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DatabaseSyncedDatabaseTableDataSynchronizationStatus
	SetInternalValue(val *DatabaseSyncedDatabaseTableDataSynchronizationStatus)
	LastSync() DatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncOutputReference
	Message() *string
	PipelineId() *string
	ProvisioningStatus() DatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusOutputReference
	ProvisioningStatusInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TriggeredUpdateStatus() DatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusOutputReference
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutContinuousUpdateStatus(value *DatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatus)
	PutFailedStatus(value *DatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatus)
	PutProvisioningStatus(value *DatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatus)
	PutTriggeredUpdateStatus(value *DatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatus)
	ResetContinuousUpdateStatus()
	ResetFailedStatus()
	ResetProvisioningStatus()
	ResetTriggeredUpdateStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference
type jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ContinuousUpdateStatus() DatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusOutputReference {
	var returns DatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusOutputReference
	_jsii_.Get(
		j,
		"continuousUpdateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ContinuousUpdateStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousUpdateStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) DetailedState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detailedState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) FailedStatus() DatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatusOutputReference {
	var returns DatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatusOutputReference
	_jsii_.Get(
		j,
		"failedStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) FailedStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) InternalValue() *DatabaseSyncedDatabaseTableDataSynchronizationStatus {
	var returns *DatabaseSyncedDatabaseTableDataSynchronizationStatus
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) LastSync() DatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncOutputReference {
	var returns DatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncOutputReference
	_jsii_.Get(
		j,
		"lastSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) PipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ProvisioningStatus() DatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusOutputReference {
	var returns DatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusOutputReference
	_jsii_.Get(
		j,
		"provisioningStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ProvisioningStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisioningStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) TriggeredUpdateStatus() DatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusOutputReference {
	var returns DatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusOutputReference
	_jsii_.Get(
		j,
		"triggeredUpdateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) TriggeredUpdateStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggeredUpdateStatusInput",
		&returns,
	)
	return returns
}


func NewDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.databaseSyncedDatabaseTable.DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference_Override(d DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.databaseSyncedDatabaseTable.DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference)SetInternalValue(val *DatabaseSyncedDatabaseTableDataSynchronizationStatus) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) PutContinuousUpdateStatus(value *DatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatus) {
	if err := d.validatePutContinuousUpdateStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContinuousUpdateStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) PutFailedStatus(value *DatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatus) {
	if err := d.validatePutFailedStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFailedStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) PutProvisioningStatus(value *DatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatus) {
	if err := d.validatePutProvisioningStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvisioningStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) PutTriggeredUpdateStatus(value *DatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatus) {
	if err := d.validatePutTriggeredUpdateStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTriggeredUpdateStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ResetContinuousUpdateStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetContinuousUpdateStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ResetFailedStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetFailedStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ResetProvisioningStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetProvisioningStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ResetTriggeredUpdateStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetTriggeredUpdateStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

