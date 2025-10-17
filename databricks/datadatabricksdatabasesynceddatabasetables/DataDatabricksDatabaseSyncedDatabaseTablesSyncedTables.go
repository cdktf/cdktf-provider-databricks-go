// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetables


type DataDatabricksDatabaseSyncedDatabaseTablesSyncedTables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/database_synced_database_tables#name DataDatabricksDatabaseSyncedDatabaseTables#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/database_synced_database_tables#database_instance_name DataDatabricksDatabaseSyncedDatabaseTables#database_instance_name}.
	DatabaseInstanceName *string `field:"optional" json:"databaseInstanceName" yaml:"databaseInstanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/database_synced_database_tables#logical_database_name DataDatabricksDatabaseSyncedDatabaseTables#logical_database_name}.
	LogicalDatabaseName *string `field:"optional" json:"logicalDatabaseName" yaml:"logicalDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/database_synced_database_tables#spec DataDatabricksDatabaseSyncedDatabaseTables#spec}.
	Spec *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpec `field:"optional" json:"spec" yaml:"spec"`
}

