// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vectorsearchindex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VectorSearchIndexConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#endpoint_name VectorSearchIndex#endpoint_name}.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#index_type VectorSearchIndex#index_type}.
	IndexType *string `field:"required" json:"indexType" yaml:"indexType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#name VectorSearchIndex#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#primary_key VectorSearchIndex#primary_key}.
	PrimaryKey *string `field:"required" json:"primaryKey" yaml:"primaryKey"`
	// delta_sync_index_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#delta_sync_index_spec VectorSearchIndex#delta_sync_index_spec}
	DeltaSyncIndexSpec *VectorSearchIndexDeltaSyncIndexSpec `field:"optional" json:"deltaSyncIndexSpec" yaml:"deltaSyncIndexSpec"`
	// direct_access_index_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#direct_access_index_spec VectorSearchIndex#direct_access_index_spec}
	DirectAccessIndexSpec *VectorSearchIndexDirectAccessIndexSpec `field:"optional" json:"directAccessIndexSpec" yaml:"directAccessIndexSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#id VectorSearchIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/vector_search_index#timeouts VectorSearchIndex#timeouts}
	Timeouts *VectorSearchIndexTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

