// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/model_serving#databricks_api_token ModelServing#databricks_api_token}.
	DatabricksApiToken *string `field:"required" json:"databricksApiToken" yaml:"databricksApiToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/model_serving#databricks_workspace_url ModelServing#databricks_workspace_url}.
	DatabricksWorkspaceUrl *string `field:"required" json:"databricksWorkspaceUrl" yaml:"databricksWorkspaceUrl"`
}
