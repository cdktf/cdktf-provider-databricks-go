// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package share


type ShareObjectPartition struct {
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.38.0/docs/resources/share#value Share#value}
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

