// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelDatabricksModelServingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/serving_endpoints#databricks_workspace_url DataDatabricksServingEndpoints#databricks_workspace_url}.
	DatabricksWorkspaceUrl *string `field:"required" json:"databricksWorkspaceUrl" yaml:"databricksWorkspaceUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/serving_endpoints#databricks_api_token DataDatabricksServingEndpoints#databricks_api_token}.
	DatabricksApiToken *string `field:"optional" json:"databricksApiToken" yaml:"databricksApiToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/serving_endpoints#databricks_api_token_plaintext DataDatabricksServingEndpoints#databricks_api_token_plaintext}.
	DatabricksApiTokenPlaintext *string `field:"optional" json:"databricksApiTokenPlaintext" yaml:"databricksApiTokenPlaintext"`
}

