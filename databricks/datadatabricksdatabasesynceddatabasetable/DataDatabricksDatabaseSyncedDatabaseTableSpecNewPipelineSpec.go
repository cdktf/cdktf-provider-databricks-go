// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetable


type DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/database_synced_database_table#budget_policy_id DataDatabricksDatabaseSyncedDatabaseTable#budget_policy_id}.
	BudgetPolicyId *string `field:"optional" json:"budgetPolicyId" yaml:"budgetPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/database_synced_database_table#storage_catalog DataDatabricksDatabaseSyncedDatabaseTable#storage_catalog}.
	StorageCatalog *string `field:"optional" json:"storageCatalog" yaml:"storageCatalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/database_synced_database_table#storage_schema DataDatabricksDatabaseSyncedDatabaseTable#storage_schema}.
	StorageSchema *string `field:"optional" json:"storageSchema" yaml:"storageSchema"`
}

