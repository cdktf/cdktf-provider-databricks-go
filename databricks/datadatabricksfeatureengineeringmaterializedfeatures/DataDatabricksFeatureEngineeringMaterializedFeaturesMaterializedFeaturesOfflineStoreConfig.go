// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringmaterializedfeatures


type DataDatabricksFeatureEngineeringMaterializedFeaturesMaterializedFeaturesOfflineStoreConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_materialized_features#catalog_name DataDatabricksFeatureEngineeringMaterializedFeatures#catalog_name}.
	CatalogName *string `field:"required" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_materialized_features#schema_name DataDatabricksFeatureEngineeringMaterializedFeatures#schema_name}.
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_materialized_features#table_name_prefix DataDatabricksFeatureEngineeringMaterializedFeatures#table_name_prefix}.
	TableNamePrefix *string `field:"required" json:"tableNamePrefix" yaml:"tableNamePrefix"`
}

