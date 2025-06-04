// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/qualitymonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/quality_monitor databricks_quality_monitor}.
type QualityMonitor interface {
	cdktf.TerraformResource
	AssetsDir() *string
	SetAssetsDir(val *string)
	AssetsDirInput() *string
	BaselineTableName() *string
	SetBaselineTableName(val *string)
	BaselineTableNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomMetrics() QualityMonitorCustomMetricsList
	CustomMetricsInput() interface{}
	DashboardId() *string
	DataClassificationConfig() QualityMonitorDataClassificationConfigList
	DataClassificationConfigInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DriftMetricsTableName() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InferenceLog() QualityMonitorInferenceLogList
	InferenceLogInput() interface{}
	LatestMonitorFailureMsg() *string
	SetLatestMonitorFailureMsg(val *string)
	LatestMonitorFailureMsgInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MonitorVersion() *string
	// The tree node.
	Node() constructs.Node
	Notifications() QualityMonitorNotificationsList
	NotificationsInput() interface{}
	OutputSchemaName() *string
	SetOutputSchemaName(val *string)
	OutputSchemaNameInput() *string
	ProfileMetricsTableName() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Schedule() QualityMonitorScheduleList
	ScheduleInput() interface{}
	SkipBuiltinDashboard() interface{}
	SetSkipBuiltinDashboard(val interface{})
	SkipBuiltinDashboardInput() interface{}
	SlicingExprs() *[]*string
	SetSlicingExprs(val *[]*string)
	SlicingExprsInput() *[]*string
	Snapshot() QualityMonitorSnapshotList
	SnapshotInput() interface{}
	Status() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeSeries() QualityMonitorTimeSeriesList
	TimeSeriesInput() interface{}
	WarehouseId() *string
	SetWarehouseId(val *string)
	WarehouseIdInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCustomMetrics(value interface{})
	PutDataClassificationConfig(value interface{})
	PutInferenceLog(value interface{})
	PutNotifications(value interface{})
	PutSchedule(value interface{})
	PutSnapshot(value interface{})
	PutTimeSeries(value interface{})
	ResetBaselineTableName()
	ResetCustomMetrics()
	ResetDataClassificationConfig()
	ResetId()
	ResetInferenceLog()
	ResetLatestMonitorFailureMsg()
	ResetNotifications()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSchedule()
	ResetSkipBuiltinDashboard()
	ResetSlicingExprs()
	ResetSnapshot()
	ResetTimeSeries()
	ResetWarehouseId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for QualityMonitor
type jsiiProxy_QualityMonitor struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_QualityMonitor) AssetsDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) AssetsDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) BaselineTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) BaselineTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) CustomMetrics() QualityMonitorCustomMetricsList {
	var returns QualityMonitorCustomMetricsList
	_jsii_.Get(
		j,
		"customMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) CustomMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) DashboardId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) DataClassificationConfig() QualityMonitorDataClassificationConfigList {
	var returns QualityMonitorDataClassificationConfigList
	_jsii_.Get(
		j,
		"dataClassificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) DataClassificationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataClassificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) DriftMetricsTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driftMetricsTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) InferenceLog() QualityMonitorInferenceLogList {
	var returns QualityMonitorInferenceLogList
	_jsii_.Get(
		j,
		"inferenceLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) InferenceLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) LatestMonitorFailureMsg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestMonitorFailureMsg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) LatestMonitorFailureMsgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestMonitorFailureMsgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) MonitorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Notifications() QualityMonitorNotificationsList {
	var returns QualityMonitorNotificationsList
	_jsii_.Get(
		j,
		"notifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) NotificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) OutputSchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) OutputSchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) ProfileMetricsTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileMetricsTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Schedule() QualityMonitorScheduleList {
	var returns QualityMonitorScheduleList
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) SkipBuiltinDashboard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipBuiltinDashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) SkipBuiltinDashboardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipBuiltinDashboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) SlicingExprs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slicingExprs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) SlicingExprsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slicingExprsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Snapshot() QualityMonitorSnapshotList {
	var returns QualityMonitorSnapshotList
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) SnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) TimeSeries() QualityMonitorTimeSeriesList {
	var returns QualityMonitorTimeSeriesList
	_jsii_.Get(
		j,
		"timeSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) TimeSeriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeSeriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitor) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/quality_monitor databricks_quality_monitor} Resource.
func NewQualityMonitor(scope constructs.Construct, id *string, config *QualityMonitorConfig) QualityMonitor {
	_init_.Initialize()

	if err := validateNewQualityMonitorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_QualityMonitor{}

	_jsii_.Create(
		"@cdktf/provider-databricks.qualityMonitor.QualityMonitor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/quality_monitor databricks_quality_monitor} Resource.
func NewQualityMonitor_Override(q QualityMonitor, scope constructs.Construct, id *string, config *QualityMonitorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.qualityMonitor.QualityMonitor",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QualityMonitor)SetAssetsDir(val *string) {
	if err := j.validateSetAssetsDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetsDir",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetBaselineTableName(val *string) {
	if err := j.validateSetBaselineTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baselineTableName",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetLatestMonitorFailureMsg(val *string) {
	if err := j.validateSetLatestMonitorFailureMsgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latestMonitorFailureMsg",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetOutputSchemaName(val *string) {
	if err := j.validateSetOutputSchemaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputSchemaName",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetSkipBuiltinDashboard(val interface{}) {
	if err := j.validateSetSkipBuiltinDashboardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipBuiltinDashboard",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetSlicingExprs(val *[]*string) {
	if err := j.validateSetSlicingExprsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slicingExprs",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_QualityMonitor)SetWarehouseId(val *string) {
	if err := j.validateSetWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

// Generates CDKTF code for importing a QualityMonitor resource upon running "cdktf plan <stack-name>".
func QualityMonitor_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateQualityMonitor_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.qualityMonitor.QualityMonitor",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func QualityMonitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQualityMonitor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.qualityMonitor.QualityMonitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QualityMonitor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQualityMonitor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.qualityMonitor.QualityMonitor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QualityMonitor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQualityMonitor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.qualityMonitor.QualityMonitor",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QualityMonitor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.qualityMonitor.QualityMonitor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (q *jsiiProxy_QualityMonitor) AddMoveTarget(moveTarget *string) {
	if err := q.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (q *jsiiProxy_QualityMonitor) AddOverride(path *string, value interface{}) {
	if err := q.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (q *jsiiProxy_QualityMonitor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := q.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (q *jsiiProxy_QualityMonitor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) MoveFromId(id *string) {
	if err := q.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveFromId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QualityMonitor) MoveTo(moveTarget *string, index interface{}) {
	if err := q.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (q *jsiiProxy_QualityMonitor) MoveToId(id *string) {
	if err := q.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveToId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QualityMonitor) OverrideLogicalId(newLogicalId *string) {
	if err := q.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QualityMonitor) PutCustomMetrics(value interface{}) {
	if err := q.validatePutCustomMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putCustomMetrics",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitor) PutDataClassificationConfig(value interface{}) {
	if err := q.validatePutDataClassificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDataClassificationConfig",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitor) PutInferenceLog(value interface{}) {
	if err := q.validatePutInferenceLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putInferenceLog",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitor) PutNotifications(value interface{}) {
	if err := q.validatePutNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putNotifications",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitor) PutSchedule(value interface{}) {
	if err := q.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSchedule",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitor) PutSnapshot(value interface{}) {
	if err := q.validatePutSnapshotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSnapshot",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitor) PutTimeSeries(value interface{}) {
	if err := q.validatePutTimeSeriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTimeSeries",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitor) ResetBaselineTableName() {
	_jsii_.InvokeVoid(
		q,
		"resetBaselineTableName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetCustomMetrics() {
	_jsii_.InvokeVoid(
		q,
		"resetCustomMetrics",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetDataClassificationConfig() {
	_jsii_.InvokeVoid(
		q,
		"resetDataClassificationConfig",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetId() {
	_jsii_.InvokeVoid(
		q,
		"resetId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetInferenceLog() {
	_jsii_.InvokeVoid(
		q,
		"resetInferenceLog",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetLatestMonitorFailureMsg() {
	_jsii_.InvokeVoid(
		q,
		"resetLatestMonitorFailureMsg",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetNotifications() {
	_jsii_.InvokeVoid(
		q,
		"resetNotifications",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetSchedule() {
	_jsii_.InvokeVoid(
		q,
		"resetSchedule",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetSkipBuiltinDashboard() {
	_jsii_.InvokeVoid(
		q,
		"resetSkipBuiltinDashboard",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetSlicingExprs() {
	_jsii_.InvokeVoid(
		q,
		"resetSlicingExprs",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetSnapshot() {
	_jsii_.InvokeVoid(
		q,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetTimeSeries() {
	_jsii_.InvokeVoid(
		q,
		"resetTimeSeries",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) ResetWarehouseId() {
	_jsii_.InvokeVoid(
		q,
		"resetWarehouseId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

