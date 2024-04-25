// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vectorsearchindex


type VectorSearchIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.41.0/docs/resources/vector_search_index#create VectorSearchIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

