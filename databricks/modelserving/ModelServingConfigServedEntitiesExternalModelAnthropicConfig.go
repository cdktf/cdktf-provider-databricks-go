// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModelAnthropicConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/model_serving#anthropic_api_key ModelServing#anthropic_api_key}.
	AnthropicApiKey *string `field:"optional" json:"anthropicApiKey" yaml:"anthropicApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/model_serving#anthropic_api_key_plaintext ModelServing#anthropic_api_key_plaintext}.
	AnthropicApiKeyPlaintext *string `field:"optional" json:"anthropicApiKeyPlaintext" yaml:"anthropicApiKeyPlaintext"`
}

