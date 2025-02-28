// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModelCohereConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/resources/model_serving#cohere_api_base ModelServing#cohere_api_base}.
	CohereApiBase *string `field:"optional" json:"cohereApiBase" yaml:"cohereApiBase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/resources/model_serving#cohere_api_key ModelServing#cohere_api_key}.
	CohereApiKey *string `field:"optional" json:"cohereApiKey" yaml:"cohereApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/resources/model_serving#cohere_api_key_plaintext ModelServing#cohere_api_key_plaintext}.
	CohereApiKeyPlaintext *string `field:"optional" json:"cohereApiKeyPlaintext" yaml:"cohereApiKeyPlaintext"`
}

