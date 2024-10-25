// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorpluginframework

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QualityMonitorPluginframeworkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#assets_dir QualityMonitorPluginframework#assets_dir}.
	AssetsDir *string `field:"required" json:"assetsDir" yaml:"assetsDir"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#output_schema_name QualityMonitorPluginframework#output_schema_name}.
	OutputSchemaName *string `field:"required" json:"outputSchemaName" yaml:"outputSchemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#table_name QualityMonitorPluginframework#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#baseline_table_name QualityMonitorPluginframework#baseline_table_name}.
	BaselineTableName *string `field:"optional" json:"baselineTableName" yaml:"baselineTableName"`
	// custom_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#custom_metrics QualityMonitorPluginframework#custom_metrics}
	CustomMetrics interface{} `field:"optional" json:"customMetrics" yaml:"customMetrics"`
	// data_classification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#data_classification_config QualityMonitorPluginframework#data_classification_config}
	DataClassificationConfig interface{} `field:"optional" json:"dataClassificationConfig" yaml:"dataClassificationConfig"`
	// inference_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#inference_log QualityMonitorPluginframework#inference_log}
	InferenceLog interface{} `field:"optional" json:"inferenceLog" yaml:"inferenceLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#latest_monitor_failure_msg QualityMonitorPluginframework#latest_monitor_failure_msg}.
	LatestMonitorFailureMsg *string `field:"optional" json:"latestMonitorFailureMsg" yaml:"latestMonitorFailureMsg"`
	// notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#notifications QualityMonitorPluginframework#notifications}
	Notifications interface{} `field:"optional" json:"notifications" yaml:"notifications"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#schedule QualityMonitorPluginframework#schedule}
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#skip_builtin_dashboard QualityMonitorPluginframework#skip_builtin_dashboard}.
	SkipBuiltinDashboard interface{} `field:"optional" json:"skipBuiltinDashboard" yaml:"skipBuiltinDashboard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#slicing_exprs QualityMonitorPluginframework#slicing_exprs}.
	SlicingExprs *[]*string `field:"optional" json:"slicingExprs" yaml:"slicingExprs"`
	// snapshot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#snapshot QualityMonitorPluginframework#snapshot}
	Snapshot interface{} `field:"optional" json:"snapshot" yaml:"snapshot"`
	// time_series block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#time_series QualityMonitorPluginframework#time_series}
	TimeSeries interface{} `field:"optional" json:"timeSeries" yaml:"timeSeries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#warehouse_id QualityMonitorPluginframework#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

