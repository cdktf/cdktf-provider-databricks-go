// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskPowerBiTaskTables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/job#catalog DataDatabricksJob#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/job#name DataDatabricksJob#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/job#schema DataDatabricksJob#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/data-sources/job#storage_mode DataDatabricksJob#storage_mode}.
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
}

