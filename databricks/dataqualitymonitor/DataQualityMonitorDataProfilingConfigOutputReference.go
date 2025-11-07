// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataqualitymonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/dataqualitymonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataQualityMonitorDataProfilingConfigOutputReference interface {
	cdktf.ComplexObject
	AssetsDir() *string
	SetAssetsDir(val *string)
	AssetsDirInput() *string
	BaselineTableName() *string
	SetBaselineTableName(val *string)
	BaselineTableNameInput() *string
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
	CustomMetrics() DataQualityMonitorDataProfilingConfigCustomMetricsList
	CustomMetricsInput() interface{}
	DashboardId() *string
	DriftMetricsTableName() *string
	EffectiveWarehouseId() *string
	// Experimental.
	Fqn() *string
	InferenceLog() DataQualityMonitorDataProfilingConfigInferenceLogOutputReference
	InferenceLogInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LatestMonitorFailureMessage() *string
	MonitoredTableName() *string
	MonitorVersion() *float64
	NotificationSettings() DataQualityMonitorDataProfilingConfigNotificationSettingsOutputReference
	NotificationSettingsInput() interface{}
	OutputSchemaId() *string
	SetOutputSchemaId(val *string)
	OutputSchemaIdInput() *string
	ProfileMetricsTableName() *string
	Schedule() DataQualityMonitorDataProfilingConfigScheduleOutputReference
	ScheduleInput() interface{}
	SkipBuiltinDashboard() interface{}
	SetSkipBuiltinDashboard(val interface{})
	SkipBuiltinDashboardInput() interface{}
	SlicingExprs() *[]*string
	SetSlicingExprs(val *[]*string)
	SlicingExprsInput() *[]*string
	Snapshot() DataQualityMonitorDataProfilingConfigSnapshotOutputReference
	SnapshotInput() interface{}
	Status() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeSeries() DataQualityMonitorDataProfilingConfigTimeSeriesOutputReference
	TimeSeriesInput() interface{}
	WarehouseId() *string
	SetWarehouseId(val *string)
	WarehouseIdInput() *string
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
	PutCustomMetrics(value interface{})
	PutInferenceLog(value *DataQualityMonitorDataProfilingConfigInferenceLog)
	PutNotificationSettings(value *DataQualityMonitorDataProfilingConfigNotificationSettings)
	PutSchedule(value *DataQualityMonitorDataProfilingConfigSchedule)
	PutSnapshot(value *DataQualityMonitorDataProfilingConfigSnapshot)
	PutTimeSeries(value *DataQualityMonitorDataProfilingConfigTimeSeries)
	ResetAssetsDir()
	ResetBaselineTableName()
	ResetCustomMetrics()
	ResetInferenceLog()
	ResetNotificationSettings()
	ResetSchedule()
	ResetSkipBuiltinDashboard()
	ResetSlicingExprs()
	ResetSnapshot()
	ResetTimeSeries()
	ResetWarehouseId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataQualityMonitorDataProfilingConfigOutputReference
type jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) AssetsDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) AssetsDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) BaselineTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) BaselineTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) CustomMetrics() DataQualityMonitorDataProfilingConfigCustomMetricsList {
	var returns DataQualityMonitorDataProfilingConfigCustomMetricsList
	_jsii_.Get(
		j,
		"customMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) CustomMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) DashboardId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) DriftMetricsTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driftMetricsTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) EffectiveWarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveWarehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) InferenceLog() DataQualityMonitorDataProfilingConfigInferenceLogOutputReference {
	var returns DataQualityMonitorDataProfilingConfigInferenceLogOutputReference
	_jsii_.Get(
		j,
		"inferenceLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) InferenceLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) LatestMonitorFailureMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestMonitorFailureMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) MonitoredTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoredTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) MonitorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) NotificationSettings() DataQualityMonitorDataProfilingConfigNotificationSettingsOutputReference {
	var returns DataQualityMonitorDataProfilingConfigNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) NotificationSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) OutputSchemaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchemaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) OutputSchemaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchemaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ProfileMetricsTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileMetricsTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) Schedule() DataQualityMonitorDataProfilingConfigScheduleOutputReference {
	var returns DataQualityMonitorDataProfilingConfigScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) SkipBuiltinDashboard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipBuiltinDashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) SkipBuiltinDashboardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipBuiltinDashboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) SlicingExprs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slicingExprs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) SlicingExprsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slicingExprsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) Snapshot() DataQualityMonitorDataProfilingConfigSnapshotOutputReference {
	var returns DataQualityMonitorDataProfilingConfigSnapshotOutputReference
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) SnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) TimeSeries() DataQualityMonitorDataProfilingConfigTimeSeriesOutputReference {
	var returns DataQualityMonitorDataProfilingConfigTimeSeriesOutputReference
	_jsii_.Get(
		j,
		"timeSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) TimeSeriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeSeriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}


func NewDataQualityMonitorDataProfilingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataQualityMonitorDataProfilingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataQualityMonitorDataProfilingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataQualityMonitor.DataQualityMonitorDataProfilingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataQualityMonitorDataProfilingConfigOutputReference_Override(d DataQualityMonitorDataProfilingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataQualityMonitor.DataQualityMonitorDataProfilingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference)SetAssetsDir(val *string) {
	if err := j.validateSetAssetsDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetsDir",
		val,
	)
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference)SetBaselineTableName(val *string) {
	if err := j.validateSetBaselineTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baselineTableName",
		val,
	)
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference)SetOutputSchemaId(val *string) {
	if err := j.validateSetOutputSchemaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputSchemaId",
		val,
	)
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference)SetSkipBuiltinDashboard(val interface{}) {
	if err := j.validateSetSkipBuiltinDashboardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipBuiltinDashboard",
		val,
	)
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference)SetSlicingExprs(val *[]*string) {
	if err := j.validateSetSlicingExprsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slicingExprs",
		val,
	)
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference)SetWarehouseId(val *string) {
	if err := j.validateSetWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) PutCustomMetrics(value interface{}) {
	if err := d.validatePutCustomMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomMetrics",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) PutInferenceLog(value *DataQualityMonitorDataProfilingConfigInferenceLog) {
	if err := d.validatePutInferenceLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInferenceLog",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) PutNotificationSettings(value *DataQualityMonitorDataProfilingConfigNotificationSettings) {
	if err := d.validatePutNotificationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotificationSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) PutSchedule(value *DataQualityMonitorDataProfilingConfigSchedule) {
	if err := d.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) PutSnapshot(value *DataQualityMonitorDataProfilingConfigSnapshot) {
	if err := d.validatePutSnapshotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSnapshot",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) PutTimeSeries(value *DataQualityMonitorDataProfilingConfigTimeSeries) {
	if err := d.validatePutTimeSeriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeSeries",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ResetAssetsDir() {
	_jsii_.InvokeVoid(
		d,
		"resetAssetsDir",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ResetBaselineTableName() {
	_jsii_.InvokeVoid(
		d,
		"resetBaselineTableName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ResetCustomMetrics() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomMetrics",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ResetInferenceLog() {
	_jsii_.InvokeVoid(
		d,
		"resetInferenceLog",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ResetNotificationSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetNotificationSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ResetSkipBuiltinDashboard() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipBuiltinDashboard",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ResetSlicingExprs() {
	_jsii_.InvokeVoid(
		d,
		"resetSlicingExprs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ResetSnapshot() {
	_jsii_.InvokeVoid(
		d,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ResetTimeSeries() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeSeries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ResetWarehouseId() {
	_jsii_.InvokeVoid(
		d,
		"resetWarehouseId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataQualityMonitorDataProfilingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

