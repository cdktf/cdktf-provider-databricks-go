// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vectorsearchindex


type VectorSearchIndexDeltaSyncIndexSpecEmbeddingSourceColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.40.0/docs/resources/vector_search_index#embedding_model_endpoint_name VectorSearchIndex#embedding_model_endpoint_name}.
	EmbeddingModelEndpointName *string `field:"optional" json:"embeddingModelEndpointName" yaml:"embeddingModelEndpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.40.0/docs/resources/vector_search_index#name VectorSearchIndex#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

