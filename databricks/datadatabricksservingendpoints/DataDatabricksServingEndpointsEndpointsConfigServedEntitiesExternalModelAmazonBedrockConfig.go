// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/serving_endpoints#aws_region DataDatabricksServingEndpoints#aws_region}.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/serving_endpoints#bedrock_provider DataDatabricksServingEndpoints#bedrock_provider}.
	BedrockProvider *string `field:"required" json:"bedrockProvider" yaml:"bedrockProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/serving_endpoints#aws_access_key_id DataDatabricksServingEndpoints#aws_access_key_id}.
	AwsAccessKeyId *string `field:"optional" json:"awsAccessKeyId" yaml:"awsAccessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/serving_endpoints#aws_access_key_id_plaintext DataDatabricksServingEndpoints#aws_access_key_id_plaintext}.
	AwsAccessKeyIdPlaintext *string `field:"optional" json:"awsAccessKeyIdPlaintext" yaml:"awsAccessKeyIdPlaintext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/serving_endpoints#aws_secret_access_key DataDatabricksServingEndpoints#aws_secret_access_key}.
	AwsSecretAccessKey *string `field:"optional" json:"awsSecretAccessKey" yaml:"awsSecretAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/serving_endpoints#aws_secret_access_key_plaintext DataDatabricksServingEndpoints#aws_secret_access_key_plaintext}.
	AwsSecretAccessKeyPlaintext *string `field:"optional" json:"awsSecretAccessKeyPlaintext" yaml:"awsSecretAccessKeyPlaintext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/serving_endpoints#instance_profile_arn DataDatabricksServingEndpoints#instance_profile_arn}.
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
}

