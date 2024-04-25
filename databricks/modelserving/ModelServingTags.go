// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.41.0/docs/resources/model_serving#key ModelServing#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.41.0/docs/resources/model_serving#value ModelServing#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

