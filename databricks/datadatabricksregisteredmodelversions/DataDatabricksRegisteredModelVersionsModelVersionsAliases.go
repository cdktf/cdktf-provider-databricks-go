// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksregisteredmodelversions


type DataDatabricksRegisteredModelVersionsModelVersionsAliases struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/registered_model_versions#alias_name DataDatabricksRegisteredModelVersions#alias_name}.
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/registered_model_versions#catalog_name DataDatabricksRegisteredModelVersions#catalog_name}.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/registered_model_versions#id DataDatabricksRegisteredModelVersions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/registered_model_versions#model_name DataDatabricksRegisteredModelVersions#model_name}.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/registered_model_versions#schema_name DataDatabricksRegisteredModelVersions#schema_name}.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/registered_model_versions#version_num DataDatabricksRegisteredModelVersions#version_num}.
	VersionNum *float64 `field:"optional" json:"versionNum" yaml:"versionNum"`
}

