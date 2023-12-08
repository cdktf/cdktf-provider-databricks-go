// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigTrafficConfigRoutes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.1/docs/resources/model_serving#served_model_name ModelServing#served_model_name}.
	ServedModelName *string `field:"required" json:"servedModelName" yaml:"servedModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.1/docs/resources/model_serving#traffic_percentage ModelServing#traffic_percentage}.
	TrafficPercentage *float64 `field:"required" json:"trafficPercentage" yaml:"trafficPercentage"`
}

