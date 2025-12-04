// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig


type FeatureEngineeringKafkaConfigSubscriptionMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/feature_engineering_kafka_config#assign FeatureEngineeringKafkaConfig#assign}.
	Assign *string `field:"optional" json:"assign" yaml:"assign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/feature_engineering_kafka_config#subscribe FeatureEngineeringKafkaConfig#subscribe}.
	Subscribe *string `field:"optional" json:"subscribe" yaml:"subscribe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/feature_engineering_kafka_config#subscribe_pattern FeatureEngineeringKafkaConfig#subscribe_pattern}.
	SubscribePattern *string `field:"optional" json:"subscribePattern" yaml:"subscribePattern"`
}

