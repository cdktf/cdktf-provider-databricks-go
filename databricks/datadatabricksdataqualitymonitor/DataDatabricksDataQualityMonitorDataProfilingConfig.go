// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitor


type DataDatabricksDataQualityMonitorDataProfilingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#output_schema_id DataDatabricksDataQualityMonitor#output_schema_id}.
	OutputSchemaId *string `field:"required" json:"outputSchemaId" yaml:"outputSchemaId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#assets_dir DataDatabricksDataQualityMonitor#assets_dir}.
	AssetsDir *string `field:"optional" json:"assetsDir" yaml:"assetsDir"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#baseline_table_name DataDatabricksDataQualityMonitor#baseline_table_name}.
	BaselineTableName *string `field:"optional" json:"baselineTableName" yaml:"baselineTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#custom_metrics DataDatabricksDataQualityMonitor#custom_metrics}.
	CustomMetrics interface{} `field:"optional" json:"customMetrics" yaml:"customMetrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#inference_log DataDatabricksDataQualityMonitor#inference_log}.
	InferenceLog *DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLog `field:"optional" json:"inferenceLog" yaml:"inferenceLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#notification_settings DataDatabricksDataQualityMonitor#notification_settings}.
	NotificationSettings *DataDatabricksDataQualityMonitorDataProfilingConfigNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#schedule DataDatabricksDataQualityMonitor#schedule}.
	Schedule *DataDatabricksDataQualityMonitorDataProfilingConfigSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#skip_builtin_dashboard DataDatabricksDataQualityMonitor#skip_builtin_dashboard}.
	SkipBuiltinDashboard interface{} `field:"optional" json:"skipBuiltinDashboard" yaml:"skipBuiltinDashboard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#slicing_exprs DataDatabricksDataQualityMonitor#slicing_exprs}.
	SlicingExprs *[]*string `field:"optional" json:"slicingExprs" yaml:"slicingExprs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#snapshot DataDatabricksDataQualityMonitor#snapshot}.
	Snapshot *DataDatabricksDataQualityMonitorDataProfilingConfigSnapshot `field:"optional" json:"snapshot" yaml:"snapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#time_series DataDatabricksDataQualityMonitor#time_series}.
	TimeSeries *DataDatabricksDataQualityMonitorDataProfilingConfigTimeSeries `field:"optional" json:"timeSeries" yaml:"timeSeries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_monitor#warehouse_id DataDatabricksDataQualityMonitor#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

