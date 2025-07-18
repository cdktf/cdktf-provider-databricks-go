// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelservingprovisionedthroughput


type ModelServingProvisionedThroughputAiGatewayGuardrailsInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving_provisioned_throughput#invalid_keywords ModelServingProvisionedThroughput#invalid_keywords}.
	InvalidKeywords *[]*string `field:"optional" json:"invalidKeywords" yaml:"invalidKeywords"`
	// pii block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving_provisioned_throughput#pii ModelServingProvisionedThroughput#pii}
	Pii *ModelServingProvisionedThroughputAiGatewayGuardrailsInputPii `field:"optional" json:"pii" yaml:"pii"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving_provisioned_throughput#safety ModelServingProvisionedThroughput#safety}.
	Safety interface{} `field:"optional" json:"safety" yaml:"safety"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving_provisioned_throughput#valid_topics ModelServingProvisionedThroughput#valid_topics}.
	ValidTopics *[]*string `field:"optional" json:"validTopics" yaml:"validTopics"`
}

