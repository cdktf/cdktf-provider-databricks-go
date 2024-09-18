// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorpluginframework


type QualityMonitorPluginframeworkInferenceLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/quality_monitor_pluginframework#granularities QualityMonitorPluginframework#granularities}.
	Granularities *[]*string `field:"required" json:"granularities" yaml:"granularities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/quality_monitor_pluginframework#model_id_col QualityMonitorPluginframework#model_id_col}.
	ModelIdCol *string `field:"required" json:"modelIdCol" yaml:"modelIdCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/quality_monitor_pluginframework#prediction_col QualityMonitorPluginframework#prediction_col}.
	PredictionCol *string `field:"required" json:"predictionCol" yaml:"predictionCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/quality_monitor_pluginframework#problem_type QualityMonitorPluginframework#problem_type}.
	ProblemType *string `field:"required" json:"problemType" yaml:"problemType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/quality_monitor_pluginframework#timestamp_col QualityMonitorPluginframework#timestamp_col}.
	TimestampCol *string `field:"required" json:"timestampCol" yaml:"timestampCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/quality_monitor_pluginframework#label_col QualityMonitorPluginframework#label_col}.
	LabelCol *string `field:"optional" json:"labelCol" yaml:"labelCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/quality_monitor_pluginframework#prediction_proba_col QualityMonitorPluginframework#prediction_proba_col}.
	PredictionProbaCol *string `field:"optional" json:"predictionProbaCol" yaml:"predictionProbaCol"`
}

