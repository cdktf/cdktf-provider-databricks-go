// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures


type DataDatabricksFeatureEngineeringFeaturesFeaturesSourceKafkaSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_features#entity_column_identifiers DataDatabricksFeatureEngineeringFeatures#entity_column_identifiers}.
	EntityColumnIdentifiers interface{} `field:"required" json:"entityColumnIdentifiers" yaml:"entityColumnIdentifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_features#name DataDatabricksFeatureEngineeringFeatures#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/feature_engineering_features#timeseries_column_identifier DataDatabricksFeatureEngineeringFeatures#timeseries_column_identifier}.
	TimeseriesColumnIdentifier *DataDatabricksFeatureEngineeringFeaturesFeaturesSourceKafkaSourceTimeseriesColumnIdentifier `field:"required" json:"timeseriesColumnIdentifier" yaml:"timeseriesColumnIdentifier"`
}

