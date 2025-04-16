// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModelPalmConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/model_serving#palm_api_key ModelServing#palm_api_key}.
	PalmApiKey *string `field:"optional" json:"palmApiKey" yaml:"palmApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/model_serving#palm_api_key_plaintext ModelServing#palm_api_key_plaintext}.
	PalmApiKeyPlaintext *string `field:"optional" json:"palmApiKeyPlaintext" yaml:"palmApiKeyPlaintext"`
}

