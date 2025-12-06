// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp


type DataDatabricksAppAppResourcesDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/app#database_name DataDatabricksApp#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/app#instance_name DataDatabricksApp#instance_name}.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/app#permission DataDatabricksApp#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

