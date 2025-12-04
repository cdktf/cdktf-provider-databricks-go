// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataqualitymonitor


type DataQualityMonitorDataProfilingConfigInferenceLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/data_quality_monitor#granularities DataQualityMonitor#granularities}.
	Granularities *[]*string `field:"required" json:"granularities" yaml:"granularities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/data_quality_monitor#model_id_column DataQualityMonitor#model_id_column}.
	ModelIdColumn *string `field:"required" json:"modelIdColumn" yaml:"modelIdColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/data_quality_monitor#prediction_column DataQualityMonitor#prediction_column}.
	PredictionColumn *string `field:"required" json:"predictionColumn" yaml:"predictionColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/data_quality_monitor#problem_type DataQualityMonitor#problem_type}.
	ProblemType *string `field:"required" json:"problemType" yaml:"problemType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/data_quality_monitor#timestamp_column DataQualityMonitor#timestamp_column}.
	TimestampColumn *string `field:"required" json:"timestampColumn" yaml:"timestampColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/data_quality_monitor#label_column DataQualityMonitor#label_column}.
	LabelColumn *string `field:"optional" json:"labelColumn" yaml:"labelColumn"`
}

