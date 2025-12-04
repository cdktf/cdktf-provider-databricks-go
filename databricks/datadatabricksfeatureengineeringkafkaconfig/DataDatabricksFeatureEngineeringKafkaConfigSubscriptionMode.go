// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfig


type DataDatabricksFeatureEngineeringKafkaConfigSubscriptionMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_kafka_config#assign DataDatabricksFeatureEngineeringKafkaConfig#assign}.
	Assign *string `field:"optional" json:"assign" yaml:"assign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_kafka_config#subscribe DataDatabricksFeatureEngineeringKafkaConfig#subscribe}.
	Subscribe *string `field:"optional" json:"subscribe" yaml:"subscribe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_kafka_config#subscribe_pattern DataDatabricksFeatureEngineeringKafkaConfig#subscribe_pattern}.
	SubscribePattern *string `field:"optional" json:"subscribePattern" yaml:"subscribePattern"`
}

