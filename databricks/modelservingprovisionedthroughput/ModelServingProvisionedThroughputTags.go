// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelservingprovisionedthroughput


type ModelServingProvisionedThroughputTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/model_serving_provisioned_throughput#key ModelServingProvisionedThroughput#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/model_serving_provisioned_throughput#value ModelServingProvisionedThroughput#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

