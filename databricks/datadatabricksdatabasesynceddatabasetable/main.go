// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetable

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTable",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInstanceName", GoGetter: "DatabaseInstanceName"},
			_jsii_.MemberProperty{JsiiProperty: "dataSynchronizationStatus", GoGetter: "DataSynchronizationStatus"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "effectiveDatabaseInstanceName", GoGetter: "EffectiveDatabaseInstanceName"},
			_jsii_.MemberProperty{JsiiProperty: "effectiveLogicalDatabaseName", GoGetter: "EffectiveLogicalDatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logicalDatabaseName", GoGetter: "LogicalDatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "unityCatalogProvisioningState", GoGetter: "UnityCatalogProvisioningState"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTable{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableConfig",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatus",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatus",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusInitialPipelineSyncProgress",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusInitialPipelineSyncProgress)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusInitialPipelineSyncProgressOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusInitialPipelineSyncProgressOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "estimatedCompletionTimeSeconds", GoGetter: "EstimatedCompletionTimeSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionCurrentlyProcessing", GoGetter: "LatestVersionCurrentlyProcessing"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningPhase", GoGetter: "ProvisioningPhase"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "syncedRowCount", GoGetter: "SyncedRowCount"},
			_jsii_.MemberProperty{JsiiProperty: "syncProgressCompletion", GoGetter: "SyncProgressCompletion"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalRowCount", GoGetter: "TotalRowCount"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusInitialPipelineSyncProgressOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "initialPipelineSyncProgress", GoGetter: "InitialPipelineSyncProgress"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastProcessedCommitVersion", GoGetter: "LastProcessedCommitVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timestamp", GoGetter: "Timestamp"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatusOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatus",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatus)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatusOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatusOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastProcessedCommitVersion", GoGetter: "LastProcessedCommitVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timestamp", GoGetter: "Timestamp"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatusOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSync",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSync)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncDeltaTableSyncInfo",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncDeltaTableSyncInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncDeltaTableSyncInfoOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncDeltaTableSyncInfoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deltaCommitTimestamp", GoGetter: "DeltaCommitTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "deltaCommitVersion", GoGetter: "DeltaCommitVersion"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncDeltaTableSyncInfoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deltaTableSyncInfo", GoGetter: "DeltaTableSyncInfo"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "syncEndTimestamp", GoGetter: "SyncEndTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "syncStartTimestamp", GoGetter: "SyncStartTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusLastSyncOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "continuousUpdateStatus", GoGetter: "ContinuousUpdateStatus"},
			_jsii_.MemberProperty{JsiiProperty: "continuousUpdateStatusInput", GoGetter: "ContinuousUpdateStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "detailedState", GoGetter: "DetailedState"},
			_jsii_.MemberProperty{JsiiProperty: "failedStatus", GoGetter: "FailedStatus"},
			_jsii_.MemberProperty{JsiiProperty: "failedStatusInput", GoGetter: "FailedStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastSync", GoGetter: "LastSync"},
			_jsii_.MemberProperty{JsiiProperty: "message", GoGetter: "Message"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineId", GoGetter: "PipelineId"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningStatus", GoGetter: "ProvisioningStatus"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningStatusInput", GoGetter: "ProvisioningStatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "putContinuousUpdateStatus", GoMethod: "PutContinuousUpdateStatus"},
			_jsii_.MemberMethod{JsiiMethod: "putFailedStatus", GoMethod: "PutFailedStatus"},
			_jsii_.MemberMethod{JsiiMethod: "putProvisioningStatus", GoMethod: "PutProvisioningStatus"},
			_jsii_.MemberMethod{JsiiMethod: "putTriggeredUpdateStatus", GoMethod: "PutTriggeredUpdateStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetContinuousUpdateStatus", GoMethod: "ResetContinuousUpdateStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailedStatus", GoMethod: "ResetFailedStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetProvisioningStatus", GoMethod: "ResetProvisioningStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetTriggeredUpdateStatus", GoMethod: "ResetTriggeredUpdateStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggeredUpdateStatus", GoGetter: "TriggeredUpdateStatus"},
			_jsii_.MemberProperty{JsiiProperty: "triggeredUpdateStatusInput", GoGetter: "TriggeredUpdateStatusInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatus",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusInitialPipelineSyncProgress",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusInitialPipelineSyncProgress)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusInitialPipelineSyncProgressOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusInitialPipelineSyncProgressOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "estimatedCompletionTimeSeconds", GoGetter: "EstimatedCompletionTimeSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionCurrentlyProcessing", GoGetter: "LatestVersionCurrentlyProcessing"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningPhase", GoGetter: "ProvisioningPhase"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "syncedRowCount", GoGetter: "SyncedRowCount"},
			_jsii_.MemberProperty{JsiiProperty: "syncProgressCompletion", GoGetter: "SyncProgressCompletion"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalRowCount", GoGetter: "TotalRowCount"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusInitialPipelineSyncProgressOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "initialPipelineSyncProgress", GoGetter: "InitialPipelineSyncProgress"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatusOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatus",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatus)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastProcessedCommitVersion", GoGetter: "LastProcessedCommitVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timestamp", GoGetter: "Timestamp"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggeredUpdateProgress", GoGetter: "TriggeredUpdateProgress"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusTriggeredUpdateProgress",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusTriggeredUpdateProgress)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusTriggeredUpdateProgressOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusTriggeredUpdateProgressOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "estimatedCompletionTimeSeconds", GoGetter: "EstimatedCompletionTimeSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionCurrentlyProcessing", GoGetter: "LatestVersionCurrentlyProcessing"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningPhase", GoGetter: "ProvisioningPhase"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "syncedRowCount", GoGetter: "SyncedRowCount"},
			_jsii_.MemberProperty{JsiiProperty: "syncProgressCompletion", GoGetter: "SyncProgressCompletion"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalRowCount", GoGetter: "TotalRowCount"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatusTriggeredUpdateProgressOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableSpec",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpec",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpecOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "budgetPolicyId", GoGetter: "BudgetPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "budgetPolicyIdInput", GoGetter: "BudgetPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetBudgetPolicyId", GoMethod: "ResetBudgetPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageCatalog", GoMethod: "ResetStorageCatalog"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageSchema", GoMethod: "ResetStorageSchema"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "storageCatalog", GoGetter: "StorageCatalog"},
			_jsii_.MemberProperty{JsiiProperty: "storageCatalogInput", GoGetter: "StorageCatalogInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageSchema", GoGetter: "StorageSchema"},
			_jsii_.MemberProperty{JsiiProperty: "storageSchemaInput", GoGetter: "StorageSchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createDatabaseObjectsIfMissing", GoGetter: "CreateDatabaseObjectsIfMissing"},
			_jsii_.MemberProperty{JsiiProperty: "createDatabaseObjectsIfMissingInput", GoGetter: "CreateDatabaseObjectsIfMissingInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "existingPipelineId", GoGetter: "ExistingPipelineId"},
			_jsii_.MemberProperty{JsiiProperty: "existingPipelineIdInput", GoGetter: "ExistingPipelineIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "newPipelineSpec", GoGetter: "NewPipelineSpec"},
			_jsii_.MemberProperty{JsiiProperty: "newPipelineSpecInput", GoGetter: "NewPipelineSpecInput"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyColumns", GoGetter: "PrimaryKeyColumns"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyColumnsInput", GoGetter: "PrimaryKeyColumnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putNewPipelineSpec", GoMethod: "PutNewPipelineSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreateDatabaseObjectsIfMissing", GoMethod: "ResetCreateDatabaseObjectsIfMissing"},
			_jsii_.MemberMethod{JsiiMethod: "resetExistingPipelineId", GoMethod: "ResetExistingPipelineId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNewPipelineSpec", GoMethod: "ResetNewPipelineSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryKeyColumns", GoMethod: "ResetPrimaryKeyColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchedulingPolicy", GoMethod: "ResetSchedulingPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceTableFullName", GoMethod: "ResetSourceTableFullName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeseriesKey", GoMethod: "ResetTimeseriesKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicy", GoGetter: "SchedulingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicyInput", GoGetter: "SchedulingPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceTableFullName", GoGetter: "SourceTableFullName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceTableFullNameInput", GoGetter: "SourceTableFullNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeseriesKey", GoGetter: "TimeseriesKey"},
			_jsii_.MemberProperty{JsiiProperty: "timeseriesKeyInput", GoGetter: "TimeseriesKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
