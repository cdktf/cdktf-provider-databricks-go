// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksdataqualitymonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference interface {
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
	CustomMetrics() DataDatabricksDataQualityMonitorDataProfilingConfigCustomMetricsList
	CustomMetricsInput() interface{}
	DashboardId() *string
	DriftMetricsTableName() *string
	EffectiveWarehouseId() *string
	// Experimental.
	Fqn() *string
	InferenceLog() DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference
	InferenceLogInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LatestMonitorFailureMessage() *string
	MonitoredTableName() *string
	MonitorVersion() *float64
	NotificationSettings() DataDatabricksDataQualityMonitorDataProfilingConfigNotificationSettingsOutputReference
	NotificationSettingsInput() interface{}
	OutputSchemaId() *string
	SetOutputSchemaId(val *string)
	OutputSchemaIdInput() *string
	ProfileMetricsTableName() *string
	Schedule() DataDatabricksDataQualityMonitorDataProfilingConfigScheduleOutputReference
	ScheduleInput() interface{}
	SkipBuiltinDashboard() interface{}
	SetSkipBuiltinDashboard(val interface{})
	SkipBuiltinDashboardInput() interface{}
	SlicingExprs() *[]*string
	SetSlicingExprs(val *[]*string)
	SlicingExprsInput() *[]*string
	Snapshot() DataDatabricksDataQualityMonitorDataProfilingConfigSnapshotOutputReference
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
	TimeSeries() DataDatabricksDataQualityMonitorDataProfilingConfigTimeSeriesOutputReference
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutCustomMetrics(value interface{})
	PutInferenceLog(value *DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLog)
	PutNotificationSettings(value *DataDatabricksDataQualityMonitorDataProfilingConfigNotificationSettings)
	PutSchedule(value *DataDatabricksDataQualityMonitorDataProfilingConfigSchedule)
	PutSnapshot(value *DataDatabricksDataQualityMonitorDataProfilingConfigSnapshot)
	PutTimeSeries(value *DataDatabricksDataQualityMonitorDataProfilingConfigTimeSeries)
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
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference
type jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) AssetsDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) AssetsDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) BaselineTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) BaselineTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) CustomMetrics() DataDatabricksDataQualityMonitorDataProfilingConfigCustomMetricsList {
	var returns DataDatabricksDataQualityMonitorDataProfilingConfigCustomMetricsList
	_jsii_.Get(
		j,
		"customMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) CustomMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) DashboardId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) DriftMetricsTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driftMetricsTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) EffectiveWarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveWarehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) InferenceLog() DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference {
	var returns DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLogOutputReference
	_jsii_.Get(
		j,
		"inferenceLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) InferenceLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) LatestMonitorFailureMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestMonitorFailureMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) MonitoredTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoredTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) MonitorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) NotificationSettings() DataDatabricksDataQualityMonitorDataProfilingConfigNotificationSettingsOutputReference {
	var returns DataDatabricksDataQualityMonitorDataProfilingConfigNotificationSettingsOutputReference
	_jsii_.Get(
		j,
		"notificationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) NotificationSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) OutputSchemaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchemaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) OutputSchemaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchemaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ProfileMetricsTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileMetricsTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) Schedule() DataDatabricksDataQualityMonitorDataProfilingConfigScheduleOutputReference {
	var returns DataDatabricksDataQualityMonitorDataProfilingConfigScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) SkipBuiltinDashboard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipBuiltinDashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) SkipBuiltinDashboardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipBuiltinDashboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) SlicingExprs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slicingExprs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) SlicingExprsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slicingExprsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) Snapshot() DataDatabricksDataQualityMonitorDataProfilingConfigSnapshotOutputReference {
	var returns DataDatabricksDataQualityMonitorDataProfilingConfigSnapshotOutputReference
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) SnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) TimeSeries() DataDatabricksDataQualityMonitorDataProfilingConfigTimeSeriesOutputReference {
	var returns DataDatabricksDataQualityMonitorDataProfilingConfigTimeSeriesOutputReference
	_jsii_.Get(
		j,
		"timeSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) TimeSeriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeSeriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksDataQualityMonitorDataProfilingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksDataQualityMonitorDataProfilingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDataQualityMonitor.DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksDataQualityMonitorDataProfilingConfigOutputReference_Override(d DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDataQualityMonitor.DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference)SetAssetsDir(val *string) {
	if err := j.validateSetAssetsDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetsDir",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference)SetBaselineTableName(val *string) {
	if err := j.validateSetBaselineTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baselineTableName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference)SetOutputSchemaId(val *string) {
	if err := j.validateSetOutputSchemaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputSchemaId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference)SetSkipBuiltinDashboard(val interface{}) {
	if err := j.validateSetSkipBuiltinDashboardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipBuiltinDashboard",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference)SetSlicingExprs(val *[]*string) {
	if err := j.validateSetSlicingExprsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slicingExprs",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference)SetWarehouseId(val *string) {
	if err := j.validateSetWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) PutCustomMetrics(value interface{}) {
	if err := d.validatePutCustomMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomMetrics",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) PutInferenceLog(value *DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLog) {
	if err := d.validatePutInferenceLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInferenceLog",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) PutNotificationSettings(value *DataDatabricksDataQualityMonitorDataProfilingConfigNotificationSettings) {
	if err := d.validatePutNotificationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotificationSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) PutSchedule(value *DataDatabricksDataQualityMonitorDataProfilingConfigSchedule) {
	if err := d.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) PutSnapshot(value *DataDatabricksDataQualityMonitorDataProfilingConfigSnapshot) {
	if err := d.validatePutSnapshotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSnapshot",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) PutTimeSeries(value *DataDatabricksDataQualityMonitorDataProfilingConfigTimeSeries) {
	if err := d.validatePutTimeSeriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeSeries",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ResetAssetsDir() {
	_jsii_.InvokeVoid(
		d,
		"resetAssetsDir",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ResetBaselineTableName() {
	_jsii_.InvokeVoid(
		d,
		"resetBaselineTableName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ResetCustomMetrics() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomMetrics",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ResetInferenceLog() {
	_jsii_.InvokeVoid(
		d,
		"resetInferenceLog",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ResetNotificationSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetNotificationSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ResetSkipBuiltinDashboard() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipBuiltinDashboard",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ResetSlicingExprs() {
	_jsii_.InvokeVoid(
		d,
		"resetSlicingExprs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ResetSnapshot() {
	_jsii_.InvokeVoid(
		d,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ResetTimeSeries() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeSeries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ResetWarehouseId() {
	_jsii_.InvokeVoid(
		d,
		"resetWarehouseId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorDataProfilingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

