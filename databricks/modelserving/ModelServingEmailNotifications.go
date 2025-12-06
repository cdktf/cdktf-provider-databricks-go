// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingEmailNotifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/model_serving#on_update_failure ModelServing#on_update_failure}.
	OnUpdateFailure *[]*string `field:"optional" json:"onUpdateFailure" yaml:"onUpdateFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/model_serving#on_update_success ModelServing#on_update_success}.
	OnUpdateSuccess *[]*string `field:"optional" json:"onUpdateSuccess" yaml:"onUpdateSuccess"`
}

