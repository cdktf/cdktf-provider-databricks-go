// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package featureengineeringmaterializedfeature


type FeatureEngineeringMaterializedFeatureOnlineStoreConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/feature_engineering_materialized_feature#catalog_name FeatureEngineeringMaterializedFeature#catalog_name}.
	CatalogName *string `field:"required" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/feature_engineering_materialized_feature#online_store_name FeatureEngineeringMaterializedFeature#online_store_name}.
	OnlineStoreName *string `field:"required" json:"onlineStoreName" yaml:"onlineStoreName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/feature_engineering_materialized_feature#schema_name FeatureEngineeringMaterializedFeature#schema_name}.
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/feature_engineering_materialized_feature#table_name_prefix FeatureEngineeringMaterializedFeature#table_name_prefix}.
	TableNamePrefix *string `field:"required" json:"tableNamePrefix" yaml:"tableNamePrefix"`
}

