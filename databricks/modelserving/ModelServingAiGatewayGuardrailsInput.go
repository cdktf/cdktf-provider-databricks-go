// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingAiGatewayGuardrailsInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/resources/model_serving#invalid_keywords ModelServing#invalid_keywords}.
	InvalidKeywords *[]*string `field:"optional" json:"invalidKeywords" yaml:"invalidKeywords"`
	// pii block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/resources/model_serving#pii ModelServing#pii}
	Pii *ModelServingAiGatewayGuardrailsInputPii `field:"optional" json:"pii" yaml:"pii"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/resources/model_serving#safety ModelServing#safety}.
	Safety interface{} `field:"optional" json:"safety" yaml:"safety"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.59.0/docs/resources/model_serving#valid_topics ModelServing#valid_topics}.
	ValidTopics *[]*string `field:"optional" json:"validTopics" yaml:"validTopics"`
}

