// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasedatabasecatalog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksDatabaseDatabaseCatalogConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/database_database_catalog#database_instance_name DataDatabricksDatabaseDatabaseCatalog#database_instance_name}.
	DatabaseInstanceName *string `field:"required" json:"databaseInstanceName" yaml:"databaseInstanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/database_database_catalog#database_name DataDatabricksDatabaseDatabaseCatalog#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/database_database_catalog#name DataDatabricksDatabaseDatabaseCatalog#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/database_database_catalog#create_database_if_not_exists DataDatabricksDatabaseDatabaseCatalog#create_database_if_not_exists}.
	CreateDatabaseIfNotExists interface{} `field:"optional" json:"createDatabaseIfNotExists" yaml:"createDatabaseIfNotExists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/database_database_catalog#workspace_id DataDatabricksDatabaseDatabaseCatalog#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

