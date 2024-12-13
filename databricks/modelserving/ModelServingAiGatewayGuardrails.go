// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingAiGatewayGuardrails struct {
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/model_serving#input ModelServing#input}
	Input *ModelServingAiGatewayGuardrailsInput `field:"optional" json:"input" yaml:"input"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/model_serving#output ModelServing#output}
	Output *ModelServingAiGatewayGuardrailsOutput `field:"optional" json:"output" yaml:"output"`
}

