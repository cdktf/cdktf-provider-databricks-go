// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vectorsearchindex


type VectorSearchIndexDeltaSyncIndexSpec struct {
	// embedding_source_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#embedding_source_columns VectorSearchIndex#embedding_source_columns}
	EmbeddingSourceColumns interface{} `field:"optional" json:"embeddingSourceColumns" yaml:"embeddingSourceColumns"`
	// embedding_vector_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#embedding_vector_columns VectorSearchIndex#embedding_vector_columns}
	EmbeddingVectorColumns interface{} `field:"optional" json:"embeddingVectorColumns" yaml:"embeddingVectorColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#embedding_writeback_table VectorSearchIndex#embedding_writeback_table}.
	EmbeddingWritebackTable *string `field:"optional" json:"embeddingWritebackTable" yaml:"embeddingWritebackTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#pipeline_type VectorSearchIndex#pipeline_type}.
	PipelineType *string `field:"optional" json:"pipelineType" yaml:"pipelineType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#source_table VectorSearchIndex#source_table}.
	SourceTable *string `field:"optional" json:"sourceTable" yaml:"sourceTable"`
}

