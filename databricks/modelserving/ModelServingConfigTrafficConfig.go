// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigTrafficConfig struct {
	// routes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/model_serving#routes ModelServing#routes}
	Routes interface{} `field:"optional" json:"routes" yaml:"routes"`
}

