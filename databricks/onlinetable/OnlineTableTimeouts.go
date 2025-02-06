// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onlinetable


type OnlineTableTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.1/docs/resources/online_table#create OnlineTable#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

