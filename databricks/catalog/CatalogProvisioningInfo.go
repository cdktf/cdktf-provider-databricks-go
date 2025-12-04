// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package catalog


type CatalogProvisioningInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/catalog#state Catalog#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

