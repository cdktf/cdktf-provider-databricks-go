// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetable


type DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/data-sources/database_synced_database_table#continuous_update_status DataDatabricksDatabaseSyncedDatabaseTable#continuous_update_status}.
	ContinuousUpdateStatus *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatus `field:"optional" json:"continuousUpdateStatus" yaml:"continuousUpdateStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/data-sources/database_synced_database_table#failed_status DataDatabricksDatabaseSyncedDatabaseTable#failed_status}.
	FailedStatus *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatus `field:"optional" json:"failedStatus" yaml:"failedStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/data-sources/database_synced_database_table#provisioning_status DataDatabricksDatabaseSyncedDatabaseTable#provisioning_status}.
	ProvisioningStatus *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatus `field:"optional" json:"provisioningStatus" yaml:"provisioningStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/data-sources/database_synced_database_table#triggered_update_status DataDatabricksDatabaseSyncedDatabaseTable#triggered_update_status}.
	TriggeredUpdateStatus *DataDatabricksDatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatus `field:"optional" json:"triggeredUpdateStatus" yaml:"triggeredUpdateStatus"`
}

