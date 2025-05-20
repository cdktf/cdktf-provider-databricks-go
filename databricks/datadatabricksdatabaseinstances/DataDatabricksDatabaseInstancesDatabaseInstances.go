// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabaseinstances


type DataDatabricksDatabaseInstancesDatabaseInstances struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/data-sources/database_instances#name DataDatabricksDatabaseInstances#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/data-sources/database_instances#admin_password DataDatabricksDatabaseInstances#admin_password}.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/data-sources/database_instances#admin_rolename DataDatabricksDatabaseInstances#admin_rolename}.
	AdminRolename *string `field:"optional" json:"adminRolename" yaml:"adminRolename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/data-sources/database_instances#capacity DataDatabricksDatabaseInstances#capacity}.
	Capacity *string `field:"optional" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.80.0/docs/data-sources/database_instances#stopped DataDatabricksDatabaseInstances#stopped}.
	Stopped interface{} `field:"optional" json:"stopped" yaml:"stopped"`
}

