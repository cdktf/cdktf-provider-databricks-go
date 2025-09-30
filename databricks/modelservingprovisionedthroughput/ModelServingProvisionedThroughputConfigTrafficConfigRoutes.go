// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelservingprovisionedthroughput


type ModelServingProvisionedThroughputConfigTrafficConfigRoutes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/model_serving_provisioned_throughput#traffic_percentage ModelServingProvisionedThroughput#traffic_percentage}.
	TrafficPercentage *float64 `field:"required" json:"trafficPercentage" yaml:"trafficPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/model_serving_provisioned_throughput#served_entity_name ModelServingProvisionedThroughput#served_entity_name}.
	ServedEntityName *string `field:"optional" json:"servedEntityName" yaml:"servedEntityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/model_serving_provisioned_throughput#served_model_name ModelServingProvisionedThroughput#served_model_name}.
	ServedModelName *string `field:"optional" json:"servedModelName" yaml:"servedModelName"`
}

