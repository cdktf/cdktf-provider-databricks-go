// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelservingprovisionedthroughput


type ModelServingProvisionedThroughputAiGateway struct {
	// fallback_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/model_serving_provisioned_throughput#fallback_config ModelServingProvisionedThroughput#fallback_config}
	FallbackConfig *ModelServingProvisionedThroughputAiGatewayFallbackConfig `field:"optional" json:"fallbackConfig" yaml:"fallbackConfig"`
	// guardrails block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/model_serving_provisioned_throughput#guardrails ModelServingProvisionedThroughput#guardrails}
	Guardrails *ModelServingProvisionedThroughputAiGatewayGuardrails `field:"optional" json:"guardrails" yaml:"guardrails"`
	// inference_table_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/model_serving_provisioned_throughput#inference_table_config ModelServingProvisionedThroughput#inference_table_config}
	InferenceTableConfig *ModelServingProvisionedThroughputAiGatewayInferenceTableConfig `field:"optional" json:"inferenceTableConfig" yaml:"inferenceTableConfig"`
	// rate_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/model_serving_provisioned_throughput#rate_limits ModelServingProvisionedThroughput#rate_limits}
	RateLimits interface{} `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// usage_tracking_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/model_serving_provisioned_throughput#usage_tracking_config ModelServingProvisionedThroughput#usage_tracking_config}
	UsageTrackingConfig *ModelServingProvisionedThroughputAiGatewayUsageTrackingConfig `field:"optional" json:"usageTrackingConfig" yaml:"usageTrackingConfig"`
}

