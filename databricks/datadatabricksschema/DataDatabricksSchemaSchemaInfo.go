// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksschema


type DataDatabricksSchemaSchemaInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#browse_only DataDatabricksSchema#browse_only}.
	BrowseOnly interface{} `field:"optional" json:"browseOnly" yaml:"browseOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#catalog_name DataDatabricksSchema#catalog_name}.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#catalog_type DataDatabricksSchema#catalog_type}.
	CatalogType *string `field:"optional" json:"catalogType" yaml:"catalogType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#comment DataDatabricksSchema#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#created_at DataDatabricksSchema#created_at}.
	CreatedAt *float64 `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#created_by DataDatabricksSchema#created_by}.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// effective_predictive_optimization_flag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#effective_predictive_optimization_flag DataDatabricksSchema#effective_predictive_optimization_flag}
	EffectivePredictiveOptimizationFlag *DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlag `field:"optional" json:"effectivePredictiveOptimizationFlag" yaml:"effectivePredictiveOptimizationFlag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#enable_predictive_optimization DataDatabricksSchema#enable_predictive_optimization}.
	EnablePredictiveOptimization *string `field:"optional" json:"enablePredictiveOptimization" yaml:"enablePredictiveOptimization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#full_name DataDatabricksSchema#full_name}.
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#metastore_id DataDatabricksSchema#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#name DataDatabricksSchema#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#owner DataDatabricksSchema#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#properties DataDatabricksSchema#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#schema_id DataDatabricksSchema#schema_id}.
	SchemaId *string `field:"optional" json:"schemaId" yaml:"schemaId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#storage_location DataDatabricksSchema#storage_location}.
	StorageLocation *string `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#storage_root DataDatabricksSchema#storage_root}.
	StorageRoot *string `field:"optional" json:"storageRoot" yaml:"storageRoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#updated_at DataDatabricksSchema#updated_at}.
	UpdatedAt *float64 `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/schema#updated_by DataDatabricksSchema#updated_by}.
	UpdatedBy *string `field:"optional" json:"updatedBy" yaml:"updatedBy"`
}

