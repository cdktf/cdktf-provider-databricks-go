// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualityrefreshes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksDataQualityRefreshesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_refreshes#object_id DataDatabricksDataQualityRefreshes#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_refreshes#object_type DataDatabricksDataQualityRefreshes#object_type}.
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/data_quality_refreshes#page_size DataDatabricksDataQualityRefreshes#page_size}.
	PageSize *float64 `field:"optional" json:"pageSize" yaml:"pageSize"`
}

