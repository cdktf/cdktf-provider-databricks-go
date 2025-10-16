// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/model_serving#token ModelServing#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/model_serving#token_plaintext ModelServing#token_plaintext}.
	TokenPlaintext *string `field:"optional" json:"tokenPlaintext" yaml:"tokenPlaintext"`
}

