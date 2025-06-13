// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelservingprovisionedthroughput


type ModelServingProvisionedThroughputAiGatewayRateLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/model_serving_provisioned_throughput#calls ModelServingProvisionedThroughput#calls}.
	Calls *float64 `field:"required" json:"calls" yaml:"calls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/model_serving_provisioned_throughput#renewal_period ModelServingProvisionedThroughput#renewal_period}.
	RenewalPeriod *string `field:"required" json:"renewalPeriod" yaml:"renewalPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/model_serving_provisioned_throughput#key ModelServingProvisionedThroughput#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

