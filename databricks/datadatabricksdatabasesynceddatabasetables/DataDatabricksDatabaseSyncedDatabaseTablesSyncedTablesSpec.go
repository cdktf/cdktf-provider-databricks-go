// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetables


type DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/database_synced_database_tables#create_database_objects_if_missing DataDatabricksDatabaseSyncedDatabaseTables#create_database_objects_if_missing}.
	CreateDatabaseObjectsIfMissing interface{} `field:"optional" json:"createDatabaseObjectsIfMissing" yaml:"createDatabaseObjectsIfMissing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/database_synced_database_tables#existing_pipeline_id DataDatabricksDatabaseSyncedDatabaseTables#existing_pipeline_id}.
	ExistingPipelineId *string `field:"optional" json:"existingPipelineId" yaml:"existingPipelineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/database_synced_database_tables#new_pipeline_spec DataDatabricksDatabaseSyncedDatabaseTables#new_pipeline_spec}.
	NewPipelineSpec *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecNewPipelineSpec `field:"optional" json:"newPipelineSpec" yaml:"newPipelineSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/database_synced_database_tables#primary_key_columns DataDatabricksDatabaseSyncedDatabaseTables#primary_key_columns}.
	PrimaryKeyColumns *[]*string `field:"optional" json:"primaryKeyColumns" yaml:"primaryKeyColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/database_synced_database_tables#scheduling_policy DataDatabricksDatabaseSyncedDatabaseTables#scheduling_policy}.
	SchedulingPolicy *string `field:"optional" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/database_synced_database_tables#source_table_full_name DataDatabricksDatabaseSyncedDatabaseTables#source_table_full_name}.
	SourceTableFullName *string `field:"optional" json:"sourceTableFullName" yaml:"sourceTableFullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/database_synced_database_tables#timeseries_key DataDatabricksDatabaseSyncedDatabaseTables#timeseries_key}.
	TimeseriesKey *string `field:"optional" json:"timeseriesKey" yaml:"timeseriesKey"`
}

