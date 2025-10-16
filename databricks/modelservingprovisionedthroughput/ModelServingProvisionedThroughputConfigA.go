// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelservingprovisionedthroughput


type ModelServingProvisionedThroughputConfigA struct {
	// served_entities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/model_serving_provisioned_throughput#served_entities ModelServingProvisionedThroughput#served_entities}
	ServedEntities interface{} `field:"optional" json:"servedEntities" yaml:"servedEntities"`
	// traffic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/model_serving_provisioned_throughput#traffic_config ModelServingProvisionedThroughput#traffic_config}
	TrafficConfig *ModelServingProvisionedThroughputConfigTrafficConfig `field:"optional" json:"trafficConfig" yaml:"trafficConfig"`
}

