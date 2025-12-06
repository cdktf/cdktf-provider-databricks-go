// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps


type DataDatabricksAppsAppResourcesDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps#database_name DataDatabricksApps#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps#instance_name DataDatabricksApps#instance_name}.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/apps#permission DataDatabricksApps#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

