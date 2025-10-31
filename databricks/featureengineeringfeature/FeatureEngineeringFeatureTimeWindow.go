// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureTimeWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/feature_engineering_feature#continuous FeatureEngineeringFeature#continuous}.
	Continuous *FeatureEngineeringFeatureTimeWindowContinuous `field:"optional" json:"continuous" yaml:"continuous"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/feature_engineering_feature#sliding FeatureEngineeringFeature#sliding}.
	Sliding *FeatureEngineeringFeatureTimeWindowSliding `field:"optional" json:"sliding" yaml:"sliding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/feature_engineering_feature#tumbling FeatureEngineeringFeature#tumbling}.
	Tumbling *FeatureEngineeringFeatureTimeWindowTumbling `field:"optional" json:"tumbling" yaml:"tumbling"`
}

