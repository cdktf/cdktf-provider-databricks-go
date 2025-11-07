// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigTrafficConfigRoutes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/model_serving#traffic_percentage ModelServing#traffic_percentage}.
	TrafficPercentage *float64 `field:"required" json:"trafficPercentage" yaml:"trafficPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/model_serving#served_entity_name ModelServing#served_entity_name}.
	ServedEntityName *string `field:"optional" json:"servedEntityName" yaml:"servedEntityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/model_serving#served_model_name ModelServing#served_model_name}.
	ServedModelName *string `field:"optional" json:"servedModelName" yaml:"servedModelName"`
}

