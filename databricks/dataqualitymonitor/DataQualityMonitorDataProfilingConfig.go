// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataqualitymonitor


type DataQualityMonitorDataProfilingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#output_schema_id DataQualityMonitor#output_schema_id}.
	OutputSchemaId *string `field:"required" json:"outputSchemaId" yaml:"outputSchemaId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#assets_dir DataQualityMonitor#assets_dir}.
	AssetsDir *string `field:"optional" json:"assetsDir" yaml:"assetsDir"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#baseline_table_name DataQualityMonitor#baseline_table_name}.
	BaselineTableName *string `field:"optional" json:"baselineTableName" yaml:"baselineTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#custom_metrics DataQualityMonitor#custom_metrics}.
	CustomMetrics interface{} `field:"optional" json:"customMetrics" yaml:"customMetrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#inference_log DataQualityMonitor#inference_log}.
	InferenceLog *DataQualityMonitorDataProfilingConfigInferenceLog `field:"optional" json:"inferenceLog" yaml:"inferenceLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#notification_settings DataQualityMonitor#notification_settings}.
	NotificationSettings *DataQualityMonitorDataProfilingConfigNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#schedule DataQualityMonitor#schedule}.
	Schedule *DataQualityMonitorDataProfilingConfigSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#skip_builtin_dashboard DataQualityMonitor#skip_builtin_dashboard}.
	SkipBuiltinDashboard interface{} `field:"optional" json:"skipBuiltinDashboard" yaml:"skipBuiltinDashboard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#slicing_exprs DataQualityMonitor#slicing_exprs}.
	SlicingExprs *[]*string `field:"optional" json:"slicingExprs" yaml:"slicingExprs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#snapshot DataQualityMonitor#snapshot}.
	Snapshot *DataQualityMonitorDataProfilingConfigSnapshot `field:"optional" json:"snapshot" yaml:"snapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#time_series DataQualityMonitor#time_series}.
	TimeSeries *DataQualityMonitorDataProfilingConfigTimeSeries `field:"optional" json:"timeSeries" yaml:"timeSeries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/data_quality_monitor#warehouse_id DataQualityMonitor#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

