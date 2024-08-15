// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskDbtTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/data-sources/job#commands DataDatabricksJob#commands}.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/data-sources/job#catalog DataDatabricksJob#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/data-sources/job#profiles_directory DataDatabricksJob#profiles_directory}.
	ProfilesDirectory *string `field:"optional" json:"profilesDirectory" yaml:"profilesDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/data-sources/job#project_directory DataDatabricksJob#project_directory}.
	ProjectDirectory *string `field:"optional" json:"projectDirectory" yaml:"projectDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/data-sources/job#schema DataDatabricksJob#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/data-sources/job#source DataDatabricksJob#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/data-sources/job#warehouse_id DataDatabricksJob#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

