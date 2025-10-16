// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasesynceddatabasetable


type DatabaseSyncedDatabaseTableSpecNewPipelineSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/database_synced_database_table#storage_catalog DatabaseSyncedDatabaseTable#storage_catalog}.
	StorageCatalog *string `field:"optional" json:"storageCatalog" yaml:"storageCatalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/database_synced_database_table#storage_schema DatabaseSyncedDatabaseTable#storage_schema}.
	StorageSchema *string `field:"optional" json:"storageSchema" yaml:"storageSchema"`
}

