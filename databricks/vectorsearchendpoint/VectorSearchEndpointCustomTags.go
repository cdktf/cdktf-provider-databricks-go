// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vectorsearchendpoint


type VectorSearchEndpointCustomTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/resources/vector_search_endpoint#key VectorSearchEndpoint#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/resources/vector_search_endpoint#value VectorSearchEndpoint#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

