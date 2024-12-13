// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppComputeStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/app#message App#message}.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/app#state App#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

