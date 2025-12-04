// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures


type DataDatabricksFeatureEngineeringFeaturesFeaturesFunction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_features#function_type DataDatabricksFeatureEngineeringFeatures#function_type}.
	FunctionType *string `field:"required" json:"functionType" yaml:"functionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_features#extra_parameters DataDatabricksFeatureEngineeringFeatures#extra_parameters}.
	ExtraParameters interface{} `field:"optional" json:"extraParameters" yaml:"extraParameters"`
}

