// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vectorsearchindex


type VectorSearchIndexDirectAccessIndexSpecEmbeddingVectorColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.1/docs/resources/vector_search_index#embedding_dimension VectorSearchIndex#embedding_dimension}.
	EmbeddingDimension *float64 `field:"optional" json:"embeddingDimension" yaml:"embeddingDimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.1/docs/resources/vector_search_index#name VectorSearchIndex#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

