// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasesynceddatabasetable


type DatabaseSyncedDatabaseTableDataSynchronizationStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/database_synced_database_table#continuous_update_status DatabaseSyncedDatabaseTable#continuous_update_status}.
	ContinuousUpdateStatus *DatabaseSyncedDatabaseTableDataSynchronizationStatusContinuousUpdateStatus `field:"optional" json:"continuousUpdateStatus" yaml:"continuousUpdateStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/database_synced_database_table#failed_status DatabaseSyncedDatabaseTable#failed_status}.
	FailedStatus *DatabaseSyncedDatabaseTableDataSynchronizationStatusFailedStatus `field:"optional" json:"failedStatus" yaml:"failedStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/database_synced_database_table#provisioning_status DatabaseSyncedDatabaseTable#provisioning_status}.
	ProvisioningStatus *DatabaseSyncedDatabaseTableDataSynchronizationStatusProvisioningStatus `field:"optional" json:"provisioningStatus" yaml:"provisioningStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/database_synced_database_table#triggered_update_status DatabaseSyncedDatabaseTable#triggered_update_status}.
	TriggeredUpdateStatus *DatabaseSyncedDatabaseTableDataSynchronizationStatusTriggeredUpdateStatus `field:"optional" json:"triggeredUpdateStatus" yaml:"triggeredUpdateStatus"`
}

