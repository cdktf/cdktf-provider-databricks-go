// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingAiGateway struct {
	// fallback_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/model_serving#fallback_config ModelServing#fallback_config}
	FallbackConfig *ModelServingAiGatewayFallbackConfig `field:"optional" json:"fallbackConfig" yaml:"fallbackConfig"`
	// guardrails block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/model_serving#guardrails ModelServing#guardrails}
	Guardrails *ModelServingAiGatewayGuardrails `field:"optional" json:"guardrails" yaml:"guardrails"`
	// inference_table_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/model_serving#inference_table_config ModelServing#inference_table_config}
	InferenceTableConfig *ModelServingAiGatewayInferenceTableConfig `field:"optional" json:"inferenceTableConfig" yaml:"inferenceTableConfig"`
	// rate_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/model_serving#rate_limits ModelServing#rate_limits}
	RateLimits interface{} `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// usage_tracking_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/resources/model_serving#usage_tracking_config ModelServing#usage_tracking_config}
	UsageTrackingConfig *ModelServingAiGatewayUsageTrackingConfig `field:"optional" json:"usageTrackingConfig" yaml:"usageTrackingConfig"`
}

