// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureFunction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/feature_engineering_feature#function_type FeatureEngineeringFeature#function_type}.
	FunctionType *string `field:"required" json:"functionType" yaml:"functionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/feature_engineering_feature#extra_parameters FeatureEngineeringFeature#extra_parameters}.
	ExtraParameters interface{} `field:"optional" json:"extraParameters" yaml:"extraParameters"`
}

