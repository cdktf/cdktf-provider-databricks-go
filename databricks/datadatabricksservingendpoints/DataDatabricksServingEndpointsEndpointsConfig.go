// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/serving_endpoints#served_entities DataDatabricksServingEndpoints#served_entities}.
	ServedEntities interface{} `field:"optional" json:"servedEntities" yaml:"servedEntities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/serving_endpoints#served_models DataDatabricksServingEndpoints#served_models}.
	ServedModels interface{} `field:"optional" json:"servedModels" yaml:"servedModels"`
}

