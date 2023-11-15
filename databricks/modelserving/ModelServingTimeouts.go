// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.30.0/docs/resources/model_serving#create ModelServing#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.30.0/docs/resources/model_serving#update ModelServing#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

