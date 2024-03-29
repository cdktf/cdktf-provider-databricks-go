// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakehousemonitor


type LakehouseMonitorInferenceLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#granularities LakehouseMonitor#granularities}.
	Granularities *[]*string `field:"optional" json:"granularities" yaml:"granularities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#label_col LakehouseMonitor#label_col}.
	LabelCol *string `field:"optional" json:"labelCol" yaml:"labelCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#model_id_col LakehouseMonitor#model_id_col}.
	ModelIdCol *string `field:"optional" json:"modelIdCol" yaml:"modelIdCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#prediction_col LakehouseMonitor#prediction_col}.
	PredictionCol *string `field:"optional" json:"predictionCol" yaml:"predictionCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#prediction_proba_col LakehouseMonitor#prediction_proba_col}.
	PredictionProbaCol *string `field:"optional" json:"predictionProbaCol" yaml:"predictionProbaCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#problem_type LakehouseMonitor#problem_type}.
	ProblemType *string `field:"optional" json:"problemType" yaml:"problemType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#timestamp_col LakehouseMonitor#timestamp_col}.
	TimestampCol *string `field:"optional" json:"timestampCol" yaml:"timestampCol"`
}

