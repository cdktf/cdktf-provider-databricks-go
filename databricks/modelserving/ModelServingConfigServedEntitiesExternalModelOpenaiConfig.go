// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModelOpenaiConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/model_serving#microsoft_entra_client_id ModelServing#microsoft_entra_client_id}.
	MicrosoftEntraClientId *string `field:"optional" json:"microsoftEntraClientId" yaml:"microsoftEntraClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/model_serving#microsoft_entra_client_secret ModelServing#microsoft_entra_client_secret}.
	MicrosoftEntraClientSecret *string `field:"optional" json:"microsoftEntraClientSecret" yaml:"microsoftEntraClientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/model_serving#microsoft_entra_client_secret_plaintext ModelServing#microsoft_entra_client_secret_plaintext}.
	MicrosoftEntraClientSecretPlaintext *string `field:"optional" json:"microsoftEntraClientSecretPlaintext" yaml:"microsoftEntraClientSecretPlaintext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/model_serving#microsoft_entra_tenant_id ModelServing#microsoft_entra_tenant_id}.
	MicrosoftEntraTenantId *string `field:"optional" json:"microsoftEntraTenantId" yaml:"microsoftEntraTenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/model_serving#openai_api_base ModelServing#openai_api_base}.
	OpenaiApiBase *string `field:"optional" json:"openaiApiBase" yaml:"openaiApiBase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/model_serving#openai_api_key ModelServing#openai_api_key}.
	OpenaiApiKey *string `field:"optional" json:"openaiApiKey" yaml:"openaiApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/model_serving#openai_api_key_plaintext ModelServing#openai_api_key_plaintext}.
	OpenaiApiKeyPlaintext *string `field:"optional" json:"openaiApiKeyPlaintext" yaml:"openaiApiKeyPlaintext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/model_serving#openai_api_type ModelServing#openai_api_type}.
	OpenaiApiType *string `field:"optional" json:"openaiApiType" yaml:"openaiApiType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/model_serving#openai_api_version ModelServing#openai_api_version}.
	OpenaiApiVersion *string `field:"optional" json:"openaiApiVersion" yaml:"openaiApiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/model_serving#openai_deployment_name ModelServing#openai_deployment_name}.
	OpenaiDeploymentName *string `field:"optional" json:"openaiDeploymentName" yaml:"openaiDeploymentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/model_serving#openai_organization ModelServing#openai_organization}.
	OpenaiOrganization *string `field:"optional" json:"openaiOrganization" yaml:"openaiOrganization"`
}

