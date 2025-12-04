// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringmaterializedfeature


type DataDatabricksFeatureEngineeringMaterializedFeatureOnlineStoreConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_materialized_feature#catalog_name DataDatabricksFeatureEngineeringMaterializedFeature#catalog_name}.
	CatalogName *string `field:"required" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_materialized_feature#online_store_name DataDatabricksFeatureEngineeringMaterializedFeature#online_store_name}.
	OnlineStoreName *string `field:"required" json:"onlineStoreName" yaml:"onlineStoreName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_materialized_feature#schema_name DataDatabricksFeatureEngineeringMaterializedFeature#schema_name}.
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_materialized_feature#table_name_prefix DataDatabricksFeatureEngineeringMaterializedFeature#table_name_prefix}.
	TableNamePrefix *string `field:"required" json:"tableNamePrefix" yaml:"tableNamePrefix"`
}

