// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package recipient


type RecipientPropertiesKvpairs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/recipient#properties Recipient#properties}.
	Properties *map[string]*string `field:"required" json:"properties" yaml:"properties"`
}

