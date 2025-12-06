// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureFunctionExtraParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_feature#key FeatureEngineeringFeature#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_feature#value FeatureEngineeringFeature#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

