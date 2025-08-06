// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/app#name App#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/app#database App#database}.
	Database *AppResourcesDatabase `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/app#description App#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/app#job App#job}.
	Job *AppResourcesJob `field:"optional" json:"job" yaml:"job"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/app#secret App#secret}.
	Secret *AppResourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/app#serving_endpoint App#serving_endpoint}.
	ServingEndpoint *AppResourcesServingEndpoint `field:"optional" json:"servingEndpoint" yaml:"servingEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/app#sql_warehouse App#sql_warehouse}.
	SqlWarehouse *AppResourcesSqlWarehouse `field:"optional" json:"sqlWarehouse" yaml:"sqlWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/app#uc_securable App#uc_securable}.
	UcSecurable *AppResourcesUcSecurable `field:"optional" json:"ucSecurable" yaml:"ucSecurable"`
}

