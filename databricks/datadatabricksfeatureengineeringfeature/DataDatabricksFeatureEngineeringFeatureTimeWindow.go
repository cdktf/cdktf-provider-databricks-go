// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureTimeWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_feature#continuous DataDatabricksFeatureEngineeringFeature#continuous}.
	Continuous *DataDatabricksFeatureEngineeringFeatureTimeWindowContinuous `field:"optional" json:"continuous" yaml:"continuous"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_feature#sliding DataDatabricksFeatureEngineeringFeature#sliding}.
	Sliding *DataDatabricksFeatureEngineeringFeatureTimeWindowSliding `field:"optional" json:"sliding" yaml:"sliding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_feature#tumbling DataDatabricksFeatureEngineeringFeature#tumbling}.
	Tumbling *DataDatabricksFeatureEngineeringFeatureTimeWindowTumbling `field:"optional" json:"tumbling" yaml:"tumbling"`
}

