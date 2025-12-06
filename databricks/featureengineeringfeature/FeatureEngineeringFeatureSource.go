// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_feature#delta_table_source FeatureEngineeringFeature#delta_table_source}.
	DeltaTableSource *FeatureEngineeringFeatureSourceDeltaTableSource `field:"optional" json:"deltaTableSource" yaml:"deltaTableSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/feature_engineering_feature#kafka_source FeatureEngineeringFeature#kafka_source}.
	KafkaSource *FeatureEngineeringFeatureSourceKafkaSource `field:"optional" json:"kafkaSource" yaml:"kafkaSource"`
}

