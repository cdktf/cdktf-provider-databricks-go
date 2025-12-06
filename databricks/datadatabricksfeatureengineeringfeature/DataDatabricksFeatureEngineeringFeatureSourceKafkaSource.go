// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureSourceKafkaSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_feature#entity_column_identifiers DataDatabricksFeatureEngineeringFeature#entity_column_identifiers}.
	EntityColumnIdentifiers interface{} `field:"required" json:"entityColumnIdentifiers" yaml:"entityColumnIdentifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_feature#name DataDatabricksFeatureEngineeringFeature#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/feature_engineering_feature#timeseries_column_identifier DataDatabricksFeatureEngineeringFeature#timeseries_column_identifier}.
	TimeseriesColumnIdentifier *DataDatabricksFeatureEngineeringFeatureSourceKafkaSourceTimeseriesColumnIdentifier `field:"required" json:"timeseriesColumnIdentifier" yaml:"timeseriesColumnIdentifier"`
}

