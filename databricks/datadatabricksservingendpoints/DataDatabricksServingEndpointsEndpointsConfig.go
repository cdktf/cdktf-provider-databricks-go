// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsConfig struct {
	// served_entities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.60.0/docs/data-sources/serving_endpoints#served_entities DataDatabricksServingEndpoints#served_entities}
	ServedEntities interface{} `field:"optional" json:"servedEntities" yaml:"servedEntities"`
	// served_models block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.60.0/docs/data-sources/serving_endpoints#served_models DataDatabricksServingEndpoints#served_models}
	ServedModels interface{} `field:"optional" json:"servedModels" yaml:"servedModels"`
}

