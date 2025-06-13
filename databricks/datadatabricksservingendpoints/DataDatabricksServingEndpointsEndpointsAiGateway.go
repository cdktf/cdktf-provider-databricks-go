// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsAiGateway struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/serving_endpoints#fallback_config DataDatabricksServingEndpoints#fallback_config}.
	FallbackConfig interface{} `field:"optional" json:"fallbackConfig" yaml:"fallbackConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/serving_endpoints#guardrails DataDatabricksServingEndpoints#guardrails}.
	Guardrails interface{} `field:"optional" json:"guardrails" yaml:"guardrails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/serving_endpoints#inference_table_config DataDatabricksServingEndpoints#inference_table_config}.
	InferenceTableConfig interface{} `field:"optional" json:"inferenceTableConfig" yaml:"inferenceTableConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/serving_endpoints#rate_limits DataDatabricksServingEndpoints#rate_limits}.
	RateLimits interface{} `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/serving_endpoints#usage_tracking_config DataDatabricksServingEndpoints#usage_tracking_config}.
	UsageTrackingConfig interface{} `field:"optional" json:"usageTrackingConfig" yaml:"usageTrackingConfig"`
}

