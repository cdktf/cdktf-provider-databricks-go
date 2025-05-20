// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitor


type QualityMonitorInferenceLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/quality_monitor#granularities QualityMonitor#granularities}.
	Granularities *[]*string `field:"required" json:"granularities" yaml:"granularities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/quality_monitor#model_id_col QualityMonitor#model_id_col}.
	ModelIdCol *string `field:"required" json:"modelIdCol" yaml:"modelIdCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/quality_monitor#prediction_col QualityMonitor#prediction_col}.
	PredictionCol *string `field:"required" json:"predictionCol" yaml:"predictionCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/quality_monitor#problem_type QualityMonitor#problem_type}.
	ProblemType *string `field:"required" json:"problemType" yaml:"problemType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/quality_monitor#timestamp_col QualityMonitor#timestamp_col}.
	TimestampCol *string `field:"required" json:"timestampCol" yaml:"timestampCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/quality_monitor#label_col QualityMonitor#label_col}.
	LabelCol *string `field:"optional" json:"labelCol" yaml:"labelCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/quality_monitor#prediction_proba_col QualityMonitor#prediction_proba_col}.
	PredictionProbaCol *string `field:"optional" json:"predictionProbaCol" yaml:"predictionProbaCol"`
}

