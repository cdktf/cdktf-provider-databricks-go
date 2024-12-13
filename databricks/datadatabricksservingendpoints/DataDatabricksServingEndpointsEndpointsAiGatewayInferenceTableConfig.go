// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsAiGatewayInferenceTableConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/serving_endpoints#catalog_name DataDatabricksServingEndpoints#catalog_name}.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/serving_endpoints#enabled DataDatabricksServingEndpoints#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/serving_endpoints#schema_name DataDatabricksServingEndpoints#schema_name}.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/serving_endpoints#table_name_prefix DataDatabricksServingEndpoints#table_name_prefix}.
	TableNamePrefix *string `field:"optional" json:"tableNamePrefix" yaml:"tableNamePrefix"`
}

