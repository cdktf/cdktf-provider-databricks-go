// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.1/docs/resources/model_serving#name ModelServing#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.1/docs/resources/model_serving#provider ModelServing#provider}.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.1/docs/resources/model_serving#task ModelServing#task}.
	Task *string `field:"required" json:"task" yaml:"task"`
	// ai21labs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.1/docs/resources/model_serving#ai21labs_config ModelServing#ai21labs_config}
	Ai21LabsConfig *ModelServingConfigServedEntitiesExternalModelAi21LabsConfig `field:"optional" json:"ai21LabsConfig" yaml:"ai21LabsConfig"`
	// amazon_bedrock_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.1/docs/resources/model_serving#amazon_bedrock_config ModelServing#amazon_bedrock_config}
	AmazonBedrockConfig *ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfig `field:"optional" json:"amazonBedrockConfig" yaml:"amazonBedrockConfig"`
	// anthropic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.1/docs/resources/model_serving#anthropic_config ModelServing#anthropic_config}
	AnthropicConfig *ModelServingConfigServedEntitiesExternalModelAnthropicConfig `field:"optional" json:"anthropicConfig" yaml:"anthropicConfig"`
	// cohere_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.1/docs/resources/model_serving#cohere_config ModelServing#cohere_config}
	CohereConfig *ModelServingConfigServedEntitiesExternalModelCohereConfig `field:"optional" json:"cohereConfig" yaml:"cohereConfig"`
	// databricks_model_serving_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.1/docs/resources/model_serving#databricks_model_serving_config ModelServing#databricks_model_serving_config}
	DatabricksModelServingConfig *ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig `field:"optional" json:"databricksModelServingConfig" yaml:"databricksModelServingConfig"`
	// openai_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.1/docs/resources/model_serving#openai_config ModelServing#openai_config}
	OpenaiConfig *ModelServingConfigServedEntitiesExternalModelOpenaiConfig `field:"optional" json:"openaiConfig" yaml:"openaiConfig"`
	// palm_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.1/docs/resources/model_serving#palm_config ModelServing#palm_config}
	PalmConfig *ModelServingConfigServedEntitiesExternalModelPalmConfig `field:"optional" json:"palmConfig" yaml:"palmConfig"`
}

