// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/resources/model_serving#aws_region ModelServing#aws_region}.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/resources/model_serving#bedrock_provider ModelServing#bedrock_provider}.
	BedrockProvider *string `field:"required" json:"bedrockProvider" yaml:"bedrockProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/resources/model_serving#aws_access_key_id ModelServing#aws_access_key_id}.
	AwsAccessKeyId *string `field:"optional" json:"awsAccessKeyId" yaml:"awsAccessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/resources/model_serving#aws_access_key_id_plaintext ModelServing#aws_access_key_id_plaintext}.
	AwsAccessKeyIdPlaintext *string `field:"optional" json:"awsAccessKeyIdPlaintext" yaml:"awsAccessKeyIdPlaintext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/resources/model_serving#aws_secret_access_key ModelServing#aws_secret_access_key}.
	AwsSecretAccessKey *string `field:"optional" json:"awsSecretAccessKey" yaml:"awsSecretAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/resources/model_serving#aws_secret_access_key_plaintext ModelServing#aws_secret_access_key_plaintext}.
	AwsSecretAccessKeyPlaintext *string `field:"optional" json:"awsSecretAccessKeyPlaintext" yaml:"awsSecretAccessKeyPlaintext"`
}

