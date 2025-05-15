// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmlflowmodel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksMlflowModelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/mlflow_model#name DataDatabricksMlflowModel#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/mlflow_model#description DataDatabricksMlflowModel#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// latest_versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/mlflow_model#latest_versions DataDatabricksMlflowModel#latest_versions}
	LatestVersions interface{} `field:"optional" json:"latestVersions" yaml:"latestVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/mlflow_model#permission_level DataDatabricksMlflowModel#permission_level}.
	PermissionLevel *string `field:"optional" json:"permissionLevel" yaml:"permissionLevel"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/mlflow_model#tags DataDatabricksMlflowModel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/mlflow_model#user_id DataDatabricksMlflowModel#user_id}.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

