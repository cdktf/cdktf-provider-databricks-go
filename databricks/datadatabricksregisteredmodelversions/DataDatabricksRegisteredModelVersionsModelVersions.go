// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksregisteredmodelversions


type DataDatabricksRegisteredModelVersionsModelVersions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#aliases DataDatabricksRegisteredModelVersions#aliases}.
	Aliases interface{} `field:"optional" json:"aliases" yaml:"aliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#catalog_name DataDatabricksRegisteredModelVersions#catalog_name}.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#comment DataDatabricksRegisteredModelVersions#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#created_at DataDatabricksRegisteredModelVersions#created_at}.
	CreatedAt *float64 `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#created_by DataDatabricksRegisteredModelVersions#created_by}.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#id DataDatabricksRegisteredModelVersions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#metastore_id DataDatabricksRegisteredModelVersions#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#model_name DataDatabricksRegisteredModelVersions#model_name}.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#model_version_dependencies DataDatabricksRegisteredModelVersions#model_version_dependencies}.
	ModelVersionDependencies interface{} `field:"optional" json:"modelVersionDependencies" yaml:"modelVersionDependencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#run_id DataDatabricksRegisteredModelVersions#run_id}.
	RunId *string `field:"optional" json:"runId" yaml:"runId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#run_workspace_id DataDatabricksRegisteredModelVersions#run_workspace_id}.
	RunWorkspaceId *float64 `field:"optional" json:"runWorkspaceId" yaml:"runWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#schema_name DataDatabricksRegisteredModelVersions#schema_name}.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#source DataDatabricksRegisteredModelVersions#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#status DataDatabricksRegisteredModelVersions#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#storage_location DataDatabricksRegisteredModelVersions#storage_location}.
	StorageLocation *string `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#updated_at DataDatabricksRegisteredModelVersions#updated_at}.
	UpdatedAt *float64 `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#updated_by DataDatabricksRegisteredModelVersions#updated_by}.
	UpdatedBy *string `field:"optional" json:"updatedBy" yaml:"updatedBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/registered_model_versions#version DataDatabricksRegisteredModelVersions#version}.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

