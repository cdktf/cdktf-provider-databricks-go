// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureLineageContext struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_feature#job_context DataDatabricksFeatureEngineeringFeature#job_context}.
	JobContext *DataDatabricksFeatureEngineeringFeatureLineageContextJobContext `field:"optional" json:"jobContext" yaml:"jobContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_feature#notebook_id DataDatabricksFeatureEngineeringFeature#notebook_id}.
	NotebookId *float64 `field:"optional" json:"notebookId" yaml:"notebookId"`
}

