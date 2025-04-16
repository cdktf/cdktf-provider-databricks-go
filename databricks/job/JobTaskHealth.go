// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskHealth struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/job#rules Job#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

