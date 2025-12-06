// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs


type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsSubscriptionMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_kafka_configs#assign DataDatabricksFeatureEngineeringKafkaConfigs#assign}.
	Assign *string `field:"optional" json:"assign" yaml:"assign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_kafka_configs#subscribe DataDatabricksFeatureEngineeringKafkaConfigs#subscribe}.
	Subscribe *string `field:"optional" json:"subscribe" yaml:"subscribe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_kafka_configs#subscribe_pattern DataDatabricksFeatureEngineeringKafkaConfigs#subscribe_pattern}.
	SubscribePattern *string `field:"optional" json:"subscribePattern" yaml:"subscribePattern"`
}

