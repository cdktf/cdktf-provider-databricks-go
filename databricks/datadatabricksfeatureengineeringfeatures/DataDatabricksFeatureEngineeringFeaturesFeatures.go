// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures


type DataDatabricksFeatureEngineeringFeaturesFeatures struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/feature_engineering_features#full_name DataDatabricksFeatureEngineeringFeatures#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/feature_engineering_features#function DataDatabricksFeatureEngineeringFeatures#function}.
	Function *DataDatabricksFeatureEngineeringFeaturesFeaturesFunction `field:"required" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/feature_engineering_features#inputs DataDatabricksFeatureEngineeringFeatures#inputs}.
	Inputs *[]*string `field:"required" json:"inputs" yaml:"inputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/feature_engineering_features#source DataDatabricksFeatureEngineeringFeatures#source}.
	Source *DataDatabricksFeatureEngineeringFeaturesFeaturesSource `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/feature_engineering_features#time_window DataDatabricksFeatureEngineeringFeatures#time_window}.
	TimeWindow *DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindow `field:"required" json:"timeWindow" yaml:"timeWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/feature_engineering_features#description DataDatabricksFeatureEngineeringFeatures#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

