// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetable


type DataDatabricksDatabaseSyncedDatabaseTableSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/database_synced_database_table#create_database_objects_if_missing DataDatabricksDatabaseSyncedDatabaseTable#create_database_objects_if_missing}.
	CreateDatabaseObjectsIfMissing interface{} `field:"optional" json:"createDatabaseObjectsIfMissing" yaml:"createDatabaseObjectsIfMissing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/database_synced_database_table#existing_pipeline_id DataDatabricksDatabaseSyncedDatabaseTable#existing_pipeline_id}.
	ExistingPipelineId *string `field:"optional" json:"existingPipelineId" yaml:"existingPipelineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/database_synced_database_table#new_pipeline_spec DataDatabricksDatabaseSyncedDatabaseTable#new_pipeline_spec}.
	NewPipelineSpec *DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpec `field:"optional" json:"newPipelineSpec" yaml:"newPipelineSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/database_synced_database_table#primary_key_columns DataDatabricksDatabaseSyncedDatabaseTable#primary_key_columns}.
	PrimaryKeyColumns *[]*string `field:"optional" json:"primaryKeyColumns" yaml:"primaryKeyColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/database_synced_database_table#scheduling_policy DataDatabricksDatabaseSyncedDatabaseTable#scheduling_policy}.
	SchedulingPolicy *string `field:"optional" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/database_synced_database_table#source_table_full_name DataDatabricksDatabaseSyncedDatabaseTable#source_table_full_name}.
	SourceTableFullName *string `field:"optional" json:"sourceTableFullName" yaml:"sourceTableFullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/database_synced_database_table#timeseries_key DataDatabricksDatabaseSyncedDatabaseTable#timeseries_key}.
	TimeseriesKey *string `field:"optional" json:"timeseriesKey" yaml:"timeseriesKey"`
}

