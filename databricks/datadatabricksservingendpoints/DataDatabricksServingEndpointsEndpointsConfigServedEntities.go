// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsConfigServedEntities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/serving_endpoints#entity_name DataDatabricksServingEndpoints#entity_name}.
	EntityName *string `field:"optional" json:"entityName" yaml:"entityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/serving_endpoints#entity_version DataDatabricksServingEndpoints#entity_version}.
	EntityVersion *string `field:"optional" json:"entityVersion" yaml:"entityVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/serving_endpoints#external_model DataDatabricksServingEndpoints#external_model}.
	ExternalModel interface{} `field:"optional" json:"externalModel" yaml:"externalModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/serving_endpoints#foundation_model DataDatabricksServingEndpoints#foundation_model}.
	FoundationModel interface{} `field:"optional" json:"foundationModel" yaml:"foundationModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/serving_endpoints#name DataDatabricksServingEndpoints#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

