// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelCustomProviderConfigApiKeyAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/serving_endpoints#key DataDatabricksServingEndpoints#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/serving_endpoints#value DataDatabricksServingEndpoints#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/serving_endpoints#value_plaintext DataDatabricksServingEndpoints#value_plaintext}.
	ValuePlaintext *string `field:"optional" json:"valuePlaintext" yaml:"valuePlaintext"`
}

