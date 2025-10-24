// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetables


type DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/database_synced_database_tables#continuous_update_status DataDatabricksDatabaseSyncedDatabaseTables#continuous_update_status}.
	ContinuousUpdateStatus *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusContinuousUpdateStatus `field:"optional" json:"continuousUpdateStatus" yaml:"continuousUpdateStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/database_synced_database_tables#failed_status DataDatabricksDatabaseSyncedDatabaseTables#failed_status}.
	FailedStatus *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusFailedStatus `field:"optional" json:"failedStatus" yaml:"failedStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/database_synced_database_tables#provisioning_status DataDatabricksDatabaseSyncedDatabaseTables#provisioning_status}.
	ProvisioningStatus *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusProvisioningStatus `field:"optional" json:"provisioningStatus" yaml:"provisioningStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/database_synced_database_tables#triggered_update_status DataDatabricksDatabaseSyncedDatabaseTables#triggered_update_status}.
	TriggeredUpdateStatus *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusTriggeredUpdateStatus `field:"optional" json:"triggeredUpdateStatus" yaml:"triggeredUpdateStatus"`
}

