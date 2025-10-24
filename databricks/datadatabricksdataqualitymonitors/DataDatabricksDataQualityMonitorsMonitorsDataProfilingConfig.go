// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitors


type DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#output_schema_id DataDatabricksDataQualityMonitors#output_schema_id}.
	OutputSchemaId *string `field:"required" json:"outputSchemaId" yaml:"outputSchemaId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#assets_dir DataDatabricksDataQualityMonitors#assets_dir}.
	AssetsDir *string `field:"optional" json:"assetsDir" yaml:"assetsDir"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#baseline_table_name DataDatabricksDataQualityMonitors#baseline_table_name}.
	BaselineTableName *string `field:"optional" json:"baselineTableName" yaml:"baselineTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#custom_metrics DataDatabricksDataQualityMonitors#custom_metrics}.
	CustomMetrics interface{} `field:"optional" json:"customMetrics" yaml:"customMetrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#inference_log DataDatabricksDataQualityMonitors#inference_log}.
	InferenceLog *DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfigInferenceLog `field:"optional" json:"inferenceLog" yaml:"inferenceLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#notification_settings DataDatabricksDataQualityMonitors#notification_settings}.
	NotificationSettings *DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfigNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#schedule DataDatabricksDataQualityMonitors#schedule}.
	Schedule *DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfigSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#skip_builtin_dashboard DataDatabricksDataQualityMonitors#skip_builtin_dashboard}.
	SkipBuiltinDashboard interface{} `field:"optional" json:"skipBuiltinDashboard" yaml:"skipBuiltinDashboard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#slicing_exprs DataDatabricksDataQualityMonitors#slicing_exprs}.
	SlicingExprs *[]*string `field:"optional" json:"slicingExprs" yaml:"slicingExprs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#snapshot DataDatabricksDataQualityMonitors#snapshot}.
	Snapshot *DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfigSnapshot `field:"optional" json:"snapshot" yaml:"snapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#time_series DataDatabricksDataQualityMonitors#time_series}.
	TimeSeries *DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfigTimeSeries `field:"optional" json:"timeSeries" yaml:"timeSeries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/data_quality_monitors#warehouse_id DataDatabricksDataQualityMonitors#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

