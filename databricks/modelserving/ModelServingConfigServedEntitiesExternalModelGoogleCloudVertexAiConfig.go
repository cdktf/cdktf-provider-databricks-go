// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.54.0/docs/resources/model_serving#private_key ModelServing#private_key}.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.54.0/docs/resources/model_serving#private_key_plaintext ModelServing#private_key_plaintext}.
	PrivateKeyPlaintext *string `field:"optional" json:"privateKeyPlaintext" yaml:"privateKeyPlaintext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.54.0/docs/resources/model_serving#project_id ModelServing#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.54.0/docs/resources/model_serving#region ModelServing#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}
