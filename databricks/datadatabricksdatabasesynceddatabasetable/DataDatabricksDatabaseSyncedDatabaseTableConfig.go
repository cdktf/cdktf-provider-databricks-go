// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksDatabaseSyncedDatabaseTableConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/database_synced_database_table#name DataDatabricksDatabaseSyncedDatabaseTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/database_synced_database_table#database_instance_name DataDatabricksDatabaseSyncedDatabaseTable#database_instance_name}.
	DatabaseInstanceName *string `field:"optional" json:"databaseInstanceName" yaml:"databaseInstanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/database_synced_database_table#logical_database_name DataDatabricksDatabaseSyncedDatabaseTable#logical_database_name}.
	LogicalDatabaseName *string `field:"optional" json:"logicalDatabaseName" yaml:"logicalDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/data-sources/database_synced_database_table#spec DataDatabricksDatabaseSyncedDatabaseTable#spec}.
	Spec *DataDatabricksDatabaseSyncedDatabaseTableSpec `field:"optional" json:"spec" yaml:"spec"`
}

