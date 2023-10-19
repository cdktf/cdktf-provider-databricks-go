// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mount


type MountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.28.1/docs/resources/mount#default Mount#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
}

