// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitor


type DataDatabricksDataQualityMonitorDataProfilingConfigInferenceLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/data_quality_monitor#granularities DataDatabricksDataQualityMonitor#granularities}.
	Granularities *[]*string `field:"required" json:"granularities" yaml:"granularities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/data_quality_monitor#model_id_column DataDatabricksDataQualityMonitor#model_id_column}.
	ModelIdColumn *string `field:"required" json:"modelIdColumn" yaml:"modelIdColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/data_quality_monitor#prediction_column DataDatabricksDataQualityMonitor#prediction_column}.
	PredictionColumn *string `field:"required" json:"predictionColumn" yaml:"predictionColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/data_quality_monitor#problem_type DataDatabricksDataQualityMonitor#problem_type}.
	ProblemType *string `field:"required" json:"problemType" yaml:"problemType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/data_quality_monitor#timestamp_column DataDatabricksDataQualityMonitor#timestamp_column}.
	TimestampColumn *string `field:"required" json:"timestampColumn" yaml:"timestampColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/data_quality_monitor#label_column DataDatabricksDataQualityMonitor#label_column}.
	LabelColumn *string `field:"optional" json:"labelColumn" yaml:"labelColumn"`
}

