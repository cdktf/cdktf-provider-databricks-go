// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/data-sources/serving_endpoints#name DataDatabricksServingEndpoints#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/data-sources/serving_endpoints#provider DataDatabricksServingEndpoints#provider}.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/data-sources/serving_endpoints#task DataDatabricksServingEndpoints#task}.
	Task *string `field:"required" json:"task" yaml:"task"`
	// ai21labs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/data-sources/serving_endpoints#ai21labs_config DataDatabricksServingEndpoints#ai21labs_config}
	Ai21LabsConfig interface{} `field:"optional" json:"ai21LabsConfig" yaml:"ai21LabsConfig"`
	// amazon_bedrock_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/data-sources/serving_endpoints#amazon_bedrock_config DataDatabricksServingEndpoints#amazon_bedrock_config}
	AmazonBedrockConfig interface{} `field:"optional" json:"amazonBedrockConfig" yaml:"amazonBedrockConfig"`
	// anthropic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/data-sources/serving_endpoints#anthropic_config DataDatabricksServingEndpoints#anthropic_config}
	AnthropicConfig interface{} `field:"optional" json:"anthropicConfig" yaml:"anthropicConfig"`
	// cohere_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/data-sources/serving_endpoints#cohere_config DataDatabricksServingEndpoints#cohere_config}
	CohereConfig interface{} `field:"optional" json:"cohereConfig" yaml:"cohereConfig"`
	// databricks_model_serving_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/data-sources/serving_endpoints#databricks_model_serving_config DataDatabricksServingEndpoints#databricks_model_serving_config}
	DatabricksModelServingConfig interface{} `field:"optional" json:"databricksModelServingConfig" yaml:"databricksModelServingConfig"`
	// google_cloud_vertex_ai_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/data-sources/serving_endpoints#google_cloud_vertex_ai_config DataDatabricksServingEndpoints#google_cloud_vertex_ai_config}
	GoogleCloudVertexAiConfig interface{} `field:"optional" json:"googleCloudVertexAiConfig" yaml:"googleCloudVertexAiConfig"`
	// openai_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/data-sources/serving_endpoints#openai_config DataDatabricksServingEndpoints#openai_config}
	OpenaiConfig interface{} `field:"optional" json:"openaiConfig" yaml:"openaiConfig"`
	// palm_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/data-sources/serving_endpoints#palm_config DataDatabricksServingEndpoints#palm_config}
	PalmConfig interface{} `field:"optional" json:"palmConfig" yaml:"palmConfig"`
}
