// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsAiGatewayGuardrailsInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/serving_endpoints#invalid_keywords DataDatabricksServingEndpoints#invalid_keywords}.
	InvalidKeywords *[]*string `field:"optional" json:"invalidKeywords" yaml:"invalidKeywords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/serving_endpoints#pii DataDatabricksServingEndpoints#pii}.
	Pii interface{} `field:"optional" json:"pii" yaml:"pii"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/serving_endpoints#safety DataDatabricksServingEndpoints#safety}.
	Safety interface{} `field:"optional" json:"safety" yaml:"safety"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/serving_endpoints#valid_topics DataDatabricksServingEndpoints#valid_topics}.
	ValidTopics *[]*string `field:"optional" json:"validTopics" yaml:"validTopics"`
}

