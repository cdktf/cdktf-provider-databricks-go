// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorpluginframework

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/qualitymonitorpluginframework/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/quality_monitor_pluginframework databricks_quality_monitor_pluginframework}.
type QualityMonitorPluginframework interface {
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
	CustomMetrics() QualityMonitorPluginframeworkCustomMetricsList
	CustomMetricsInput() interface{}
	DashboardId() *string
	DataClassificationConfig() QualityMonitorPluginframeworkDataClassificationConfigOutputReference
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
	InferenceLog() QualityMonitorPluginframeworkInferenceLogOutputReference
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
	Notifications() QualityMonitorPluginframeworkNotificationsOutputReference
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
	Schedule() QualityMonitorPluginframeworkScheduleOutputReference
	ScheduleInput() interface{}
	SkipBuiltinDashboard() interface{}
	SetSkipBuiltinDashboard(val interface{})
	SkipBuiltinDashboardInput() interface{}
	SlicingExprs() *[]*string
	SetSlicingExprs(val *[]*string)
	SlicingExprsInput() *[]*string
	Snapshot() QualityMonitorPluginframeworkSnapshotOutputReference
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
	TimeSeries() QualityMonitorPluginframeworkTimeSeriesOutputReference
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
	PutDataClassificationConfig(value *QualityMonitorPluginframeworkDataClassificationConfig)
	PutInferenceLog(value *QualityMonitorPluginframeworkInferenceLog)
	PutNotifications(value *QualityMonitorPluginframeworkNotifications)
	PutSchedule(value *QualityMonitorPluginframeworkSchedule)
	PutSnapshot(value *QualityMonitorPluginframeworkSnapshot)
	PutTimeSeries(value *QualityMonitorPluginframeworkTimeSeries)
	ResetBaselineTableName()
	ResetCustomMetrics()
	ResetDataClassificationConfig()
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

// The jsii proxy struct for QualityMonitorPluginframework
type jsiiProxy_QualityMonitorPluginframework struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_QualityMonitorPluginframework) AssetsDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) AssetsDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) BaselineTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) BaselineTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) CustomMetrics() QualityMonitorPluginframeworkCustomMetricsList {
	var returns QualityMonitorPluginframeworkCustomMetricsList
	_jsii_.Get(
		j,
		"customMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) CustomMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) DashboardId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) DataClassificationConfig() QualityMonitorPluginframeworkDataClassificationConfigOutputReference {
	var returns QualityMonitorPluginframeworkDataClassificationConfigOutputReference
	_jsii_.Get(
		j,
		"dataClassificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) DataClassificationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataClassificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) DriftMetricsTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driftMetricsTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) InferenceLog() QualityMonitorPluginframeworkInferenceLogOutputReference {
	var returns QualityMonitorPluginframeworkInferenceLogOutputReference
	_jsii_.Get(
		j,
		"inferenceLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) InferenceLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) LatestMonitorFailureMsg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestMonitorFailureMsg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) LatestMonitorFailureMsgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestMonitorFailureMsgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) MonitorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) Notifications() QualityMonitorPluginframeworkNotificationsOutputReference {
	var returns QualityMonitorPluginframeworkNotificationsOutputReference
	_jsii_.Get(
		j,
		"notifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) NotificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) OutputSchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) OutputSchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSchemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) ProfileMetricsTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileMetricsTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) Schedule() QualityMonitorPluginframeworkScheduleOutputReference {
	var returns QualityMonitorPluginframeworkScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) SkipBuiltinDashboard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipBuiltinDashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) SkipBuiltinDashboardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipBuiltinDashboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) SlicingExprs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slicingExprs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) SlicingExprsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slicingExprsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) Snapshot() QualityMonitorPluginframeworkSnapshotOutputReference {
	var returns QualityMonitorPluginframeworkSnapshotOutputReference
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) SnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) TimeSeries() QualityMonitorPluginframeworkTimeSeriesOutputReference {
	var returns QualityMonitorPluginframeworkTimeSeriesOutputReference
	_jsii_.Get(
		j,
		"timeSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) TimeSeriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeSeriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorPluginframework) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/quality_monitor_pluginframework databricks_quality_monitor_pluginframework} Resource.
func NewQualityMonitorPluginframework(scope constructs.Construct, id *string, config *QualityMonitorPluginframeworkConfig) QualityMonitorPluginframework {
	_init_.Initialize()

	if err := validateNewQualityMonitorPluginframeworkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_QualityMonitorPluginframework{}

	_jsii_.Create(
		"@cdktf/provider-databricks.qualityMonitorPluginframework.QualityMonitorPluginframework",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/quality_monitor_pluginframework databricks_quality_monitor_pluginframework} Resource.
func NewQualityMonitorPluginframework_Override(q QualityMonitorPluginframework, scope constructs.Construct, id *string, config *QualityMonitorPluginframeworkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.qualityMonitorPluginframework.QualityMonitorPluginframework",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetAssetsDir(val *string) {
	if err := j.validateSetAssetsDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetsDir",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetBaselineTableName(val *string) {
	if err := j.validateSetBaselineTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baselineTableName",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetLatestMonitorFailureMsg(val *string) {
	if err := j.validateSetLatestMonitorFailureMsgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latestMonitorFailureMsg",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetOutputSchemaName(val *string) {
	if err := j.validateSetOutputSchemaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputSchemaName",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetSkipBuiltinDashboard(val interface{}) {
	if err := j.validateSetSkipBuiltinDashboardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipBuiltinDashboard",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetSlicingExprs(val *[]*string) {
	if err := j.validateSetSlicingExprsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slicingExprs",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorPluginframework)SetWarehouseId(val *string) {
	if err := j.validateSetWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

// Generates CDKTF code for importing a QualityMonitorPluginframework resource upon running "cdktf plan <stack-name>".
func QualityMonitorPluginframework_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateQualityMonitorPluginframework_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.qualityMonitorPluginframework.QualityMonitorPluginframework",
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
func QualityMonitorPluginframework_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQualityMonitorPluginframework_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.qualityMonitorPluginframework.QualityMonitorPluginframework",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QualityMonitorPluginframework_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQualityMonitorPluginframework_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.qualityMonitorPluginframework.QualityMonitorPluginframework",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QualityMonitorPluginframework_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQualityMonitorPluginframework_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.qualityMonitorPluginframework.QualityMonitorPluginframework",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QualityMonitorPluginframework_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.qualityMonitorPluginframework.QualityMonitorPluginframework",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframework) AddMoveTarget(moveTarget *string) {
	if err := q.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) AddOverride(path *string, value interface{}) {
	if err := q.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QualityMonitorPluginframework) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QualityMonitorPluginframework) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QualityMonitorPluginframework) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QualityMonitorPluginframework) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QualityMonitorPluginframework) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QualityMonitorPluginframework) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QualityMonitorPluginframework) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QualityMonitorPluginframework) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QualityMonitorPluginframework) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframework) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := q.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QualityMonitorPluginframework) MoveFromId(id *string) {
	if err := q.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveFromId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) MoveTo(moveTarget *string, index interface{}) {
	if err := q.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) MoveToId(id *string) {
	if err := q.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveToId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) OverrideLogicalId(newLogicalId *string) {
	if err := q.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) PutCustomMetrics(value interface{}) {
	if err := q.validatePutCustomMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putCustomMetrics",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) PutDataClassificationConfig(value *QualityMonitorPluginframeworkDataClassificationConfig) {
	if err := q.validatePutDataClassificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDataClassificationConfig",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) PutInferenceLog(value *QualityMonitorPluginframeworkInferenceLog) {
	if err := q.validatePutInferenceLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putInferenceLog",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) PutNotifications(value *QualityMonitorPluginframeworkNotifications) {
	if err := q.validatePutNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putNotifications",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) PutSchedule(value *QualityMonitorPluginframeworkSchedule) {
	if err := q.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSchedule",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) PutSnapshot(value *QualityMonitorPluginframeworkSnapshot) {
	if err := q.validatePutSnapshotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSnapshot",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) PutTimeSeries(value *QualityMonitorPluginframeworkTimeSeries) {
	if err := q.validatePutTimeSeriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTimeSeries",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetBaselineTableName() {
	_jsii_.InvokeVoid(
		q,
		"resetBaselineTableName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetCustomMetrics() {
	_jsii_.InvokeVoid(
		q,
		"resetCustomMetrics",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetDataClassificationConfig() {
	_jsii_.InvokeVoid(
		q,
		"resetDataClassificationConfig",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetInferenceLog() {
	_jsii_.InvokeVoid(
		q,
		"resetInferenceLog",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetLatestMonitorFailureMsg() {
	_jsii_.InvokeVoid(
		q,
		"resetLatestMonitorFailureMsg",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetNotifications() {
	_jsii_.InvokeVoid(
		q,
		"resetNotifications",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetSchedule() {
	_jsii_.InvokeVoid(
		q,
		"resetSchedule",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetSkipBuiltinDashboard() {
	_jsii_.InvokeVoid(
		q,
		"resetSkipBuiltinDashboard",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetSlicingExprs() {
	_jsii_.InvokeVoid(
		q,
		"resetSlicingExprs",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetSnapshot() {
	_jsii_.InvokeVoid(
		q,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetTimeSeries() {
	_jsii_.InvokeVoid(
		q,
		"resetTimeSeries",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) ResetWarehouseId() {
	_jsii_.InvokeVoid(
		q,
		"resetWarehouseId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorPluginframework) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframework) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframework) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframework) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframework) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorPluginframework) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

