// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QualityMonitorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#assets_dir QualityMonitor#assets_dir}.
	AssetsDir *string `field:"required" json:"assetsDir" yaml:"assetsDir"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#output_schema_name QualityMonitor#output_schema_name}.
	OutputSchemaName *string `field:"required" json:"outputSchemaName" yaml:"outputSchemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#table_name QualityMonitor#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#baseline_table_name QualityMonitor#baseline_table_name}.
	BaselineTableName *string `field:"optional" json:"baselineTableName" yaml:"baselineTableName"`
	// custom_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#custom_metrics QualityMonitor#custom_metrics}
	CustomMetrics interface{} `field:"optional" json:"customMetrics" yaml:"customMetrics"`
	// data_classification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#data_classification_config QualityMonitor#data_classification_config}
	DataClassificationConfig *QualityMonitorDataClassificationConfig `field:"optional" json:"dataClassificationConfig" yaml:"dataClassificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#id QualityMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// inference_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#inference_log QualityMonitor#inference_log}
	InferenceLog *QualityMonitorInferenceLog `field:"optional" json:"inferenceLog" yaml:"inferenceLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#latest_monitor_failure_msg QualityMonitor#latest_monitor_failure_msg}.
	LatestMonitorFailureMsg *string `field:"optional" json:"latestMonitorFailureMsg" yaml:"latestMonitorFailureMsg"`
	// notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#notifications QualityMonitor#notifications}
	Notifications *QualityMonitorNotifications `field:"optional" json:"notifications" yaml:"notifications"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#schedule QualityMonitor#schedule}
	Schedule *QualityMonitorSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#skip_builtin_dashboard QualityMonitor#skip_builtin_dashboard}.
	SkipBuiltinDashboard interface{} `field:"optional" json:"skipBuiltinDashboard" yaml:"skipBuiltinDashboard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#slicing_exprs QualityMonitor#slicing_exprs}.
	SlicingExprs *[]*string `field:"optional" json:"slicingExprs" yaml:"slicingExprs"`
	// snapshot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#snapshot QualityMonitor#snapshot}
	Snapshot *QualityMonitorSnapshot `field:"optional" json:"snapshot" yaml:"snapshot"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#timeouts QualityMonitor#timeouts}
	Timeouts *QualityMonitorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// time_series block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#time_series QualityMonitor#time_series}
	TimeSeries *QualityMonitorTimeSeries `field:"optional" json:"timeSeries" yaml:"timeSeries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor#warehouse_id QualityMonitor#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

