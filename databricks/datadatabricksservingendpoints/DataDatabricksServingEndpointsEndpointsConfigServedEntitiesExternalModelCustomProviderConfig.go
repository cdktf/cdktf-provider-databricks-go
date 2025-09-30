// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelCustomProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/serving_endpoints#custom_provider_url DataDatabricksServingEndpoints#custom_provider_url}.
	CustomProviderUrl *string `field:"required" json:"customProviderUrl" yaml:"customProviderUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/serving_endpoints#api_key_auth DataDatabricksServingEndpoints#api_key_auth}.
	ApiKeyAuth interface{} `field:"optional" json:"apiKeyAuth" yaml:"apiKeyAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/serving_endpoints#bearer_token_auth DataDatabricksServingEndpoints#bearer_token_auth}.
	BearerTokenAuth interface{} `field:"optional" json:"bearerTokenAuth" yaml:"bearerTokenAuth"`
}

