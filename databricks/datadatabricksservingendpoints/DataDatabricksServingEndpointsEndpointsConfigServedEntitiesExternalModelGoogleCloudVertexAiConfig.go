// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelGoogleCloudVertexAiConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/data-sources/serving_endpoints#project_id DataDatabricksServingEndpoints#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/data-sources/serving_endpoints#region DataDatabricksServingEndpoints#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/data-sources/serving_endpoints#private_key DataDatabricksServingEndpoints#private_key}.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/data-sources/serving_endpoints#private_key_plaintext DataDatabricksServingEndpoints#private_key_plaintext}.
	PrivateKeyPlaintext *string `field:"optional" json:"privateKeyPlaintext" yaml:"privateKeyPlaintext"`
}

