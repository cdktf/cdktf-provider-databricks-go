// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksregisteredmodel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksRegisteredModelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/registered_model#full_name DataDatabricksRegisteredModel#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/registered_model#include_aliases DataDatabricksRegisteredModel#include_aliases}.
	IncludeAliases interface{} `field:"optional" json:"includeAliases" yaml:"includeAliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/registered_model#include_browse DataDatabricksRegisteredModel#include_browse}.
	IncludeBrowse interface{} `field:"optional" json:"includeBrowse" yaml:"includeBrowse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/registered_model#model_info DataDatabricksRegisteredModel#model_info}.
	ModelInfo interface{} `field:"optional" json:"modelInfo" yaml:"modelInfo"`
}

