// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagpolicy


type TagPolicyValues struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/tag_policy#name TagPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

