// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsAiGatewayRateLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/serving_endpoints#calls DataDatabricksServingEndpoints#calls}.
	Calls *float64 `field:"required" json:"calls" yaml:"calls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/serving_endpoints#renewal_period DataDatabricksServingEndpoints#renewal_period}.
	RenewalPeriod *string `field:"required" json:"renewalPeriod" yaml:"renewalPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/serving_endpoints#key DataDatabricksServingEndpoints#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

