// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModelCustomProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/model_serving#custom_provider_url ModelServing#custom_provider_url}.
	CustomProviderUrl *string `field:"required" json:"customProviderUrl" yaml:"customProviderUrl"`
	// api_key_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/model_serving#api_key_auth ModelServing#api_key_auth}
	ApiKeyAuth *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigApiKeyAuth `field:"optional" json:"apiKeyAuth" yaml:"apiKeyAuth"`
	// bearer_token_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/model_serving#bearer_token_auth ModelServing#bearer_token_auth}
	BearerTokenAuth *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth `field:"optional" json:"bearerTokenAuth" yaml:"bearerTokenAuth"`
}

