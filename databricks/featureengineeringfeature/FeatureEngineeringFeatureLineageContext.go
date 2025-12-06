// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureLineageContext struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_feature#job_context FeatureEngineeringFeature#job_context}.
	JobContext *FeatureEngineeringFeatureLineageContextJobContext `field:"optional" json:"jobContext" yaml:"jobContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_feature#notebook_id FeatureEngineeringFeature#notebook_id}.
	NotebookId *float64 `field:"optional" json:"notebookId" yaml:"notebookId"`
}

