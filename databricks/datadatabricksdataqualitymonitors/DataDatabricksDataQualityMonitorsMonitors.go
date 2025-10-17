// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitors


type DataDatabricksDataQualityMonitorsMonitors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/data_quality_monitors#object_id DataDatabricksDataQualityMonitors#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/data_quality_monitors#object_type DataDatabricksDataQualityMonitors#object_type}.
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/data_quality_monitors#anomaly_detection_config DataDatabricksDataQualityMonitors#anomaly_detection_config}.
	AnomalyDetectionConfig *DataDatabricksDataQualityMonitorsMonitorsAnomalyDetectionConfig `field:"optional" json:"anomalyDetectionConfig" yaml:"anomalyDetectionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/data_quality_monitors#data_profiling_config DataDatabricksDataQualityMonitors#data_profiling_config}.
	DataProfilingConfig *DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfig `field:"optional" json:"dataProfilingConfig" yaml:"dataProfilingConfig"`
}

