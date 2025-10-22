// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasedatabasecatalogs


type DataDatabricksDatabaseDatabaseCatalogsDatabaseCatalogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/database_database_catalogs#database_instance_name DataDatabricksDatabaseDatabaseCatalogs#database_instance_name}.
	DatabaseInstanceName *string `field:"required" json:"databaseInstanceName" yaml:"databaseInstanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/database_database_catalogs#database_name DataDatabricksDatabaseDatabaseCatalogs#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/database_database_catalogs#name DataDatabricksDatabaseDatabaseCatalogs#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/database_database_catalogs#create_database_if_not_exists DataDatabricksDatabaseDatabaseCatalogs#create_database_if_not_exists}.
	CreateDatabaseIfNotExists interface{} `field:"optional" json:"createDatabaseIfNotExists" yaml:"createDatabaseIfNotExists"`
}

