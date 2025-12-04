// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureTimeWindowSliding struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/feature_engineering_feature#slide_duration FeatureEngineeringFeature#slide_duration}.
	SlideDuration *string `field:"required" json:"slideDuration" yaml:"slideDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/feature_engineering_feature#window_duration FeatureEngineeringFeature#window_duration}.
	WindowDuration *string `field:"required" json:"windowDuration" yaml:"windowDuration"`
}

