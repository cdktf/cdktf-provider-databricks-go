// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures


type DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_features#continuous DataDatabricksFeatureEngineeringFeatures#continuous}.
	Continuous *DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowContinuous `field:"optional" json:"continuous" yaml:"continuous"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_features#sliding DataDatabricksFeatureEngineeringFeatures#sliding}.
	Sliding *DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowSliding `field:"optional" json:"sliding" yaml:"sliding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_features#tumbling DataDatabricksFeatureEngineeringFeatures#tumbling}.
	Tumbling *DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowTumbling `field:"optional" json:"tumbling" yaml:"tumbling"`
}

