// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vectorsearchindex


type VectorSearchIndexDirectAccessIndexSpecEmbeddingSourceColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/vector_search_index#embedding_model_endpoint_name VectorSearchIndex#embedding_model_endpoint_name}.
	EmbeddingModelEndpointName *string `field:"optional" json:"embeddingModelEndpointName" yaml:"embeddingModelEndpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/vector_search_index#model_endpoint_name_for_query VectorSearchIndex#model_endpoint_name_for_query}.
	ModelEndpointNameForQuery *string `field:"optional" json:"modelEndpointNameForQuery" yaml:"modelEndpointNameForQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/resources/vector_search_index#name VectorSearchIndex#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

