// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabaseinstances


type DataDatabricksDatabaseInstancesDatabaseInstances struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/database_instances#name DataDatabricksDatabaseInstances#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/database_instances#capacity DataDatabricksDatabaseInstances#capacity}.
	Capacity *string `field:"optional" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/database_instances#enable_pg_native_login DataDatabricksDatabaseInstances#enable_pg_native_login}.
	EnablePgNativeLogin interface{} `field:"optional" json:"enablePgNativeLogin" yaml:"enablePgNativeLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/database_instances#enable_readable_secondaries DataDatabricksDatabaseInstances#enable_readable_secondaries}.
	EnableReadableSecondaries interface{} `field:"optional" json:"enableReadableSecondaries" yaml:"enableReadableSecondaries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/database_instances#node_count DataDatabricksDatabaseInstances#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/database_instances#parent_instance_ref DataDatabricksDatabaseInstances#parent_instance_ref}.
	ParentInstanceRef *DataDatabricksDatabaseInstancesDatabaseInstancesParentInstanceRef `field:"optional" json:"parentInstanceRef" yaml:"parentInstanceRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/database_instances#retention_window_in_days DataDatabricksDatabaseInstances#retention_window_in_days}.
	RetentionWindowInDays *float64 `field:"optional" json:"retentionWindowInDays" yaml:"retentionWindowInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/database_instances#stopped DataDatabricksDatabaseInstances#stopped}.
	Stopped interface{} `field:"optional" json:"stopped" yaml:"stopped"`
}

