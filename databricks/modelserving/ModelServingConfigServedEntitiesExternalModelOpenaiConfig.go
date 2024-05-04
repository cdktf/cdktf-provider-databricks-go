// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModelOpenaiConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/model_serving#openai_api_key ModelServing#openai_api_key}.
	OpenaiApiKey *string `field:"required" json:"openaiApiKey" yaml:"openaiApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/model_serving#openai_api_base ModelServing#openai_api_base}.
	OpenaiApiBase *string `field:"optional" json:"openaiApiBase" yaml:"openaiApiBase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/model_serving#openai_api_type ModelServing#openai_api_type}.
	OpenaiApiType *string `field:"optional" json:"openaiApiType" yaml:"openaiApiType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/model_serving#openai_api_version ModelServing#openai_api_version}.
	OpenaiApiVersion *string `field:"optional" json:"openaiApiVersion" yaml:"openaiApiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/model_serving#openai_deployment_name ModelServing#openai_deployment_name}.
	OpenaiDeploymentName *string `field:"optional" json:"openaiDeploymentName" yaml:"openaiDeploymentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/model_serving#openai_organization ModelServing#openai_organization}.
	OpenaiOrganization *string `field:"optional" json:"openaiOrganization" yaml:"openaiOrganization"`
}

