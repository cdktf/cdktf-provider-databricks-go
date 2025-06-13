// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppResourcesUcSecurable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/app#permission App#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/app#securable_full_name App#securable_full_name}.
	SecurableFullName *string `field:"required" json:"securableFullName" yaml:"securableFullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/app#securable_type App#securable_type}.
	SecurableType *string `field:"required" json:"securableType" yaml:"securableType"`
}

