// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModelCohereConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/model_serving#cohere_api_key ModelServing#cohere_api_key}.
	CohereApiKey *string `field:"required" json:"cohereApiKey" yaml:"cohereApiKey"`
}

