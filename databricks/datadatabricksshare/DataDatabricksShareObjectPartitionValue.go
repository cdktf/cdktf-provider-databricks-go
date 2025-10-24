// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksshare


type DataDatabricksShareObjectPartitionValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#name DataDatabricksShare#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#op DataDatabricksShare#op}.
	Op *string `field:"optional" json:"op" yaml:"op"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#recipient_property_key DataDatabricksShare#recipient_property_key}.
	RecipientPropertyKey *string `field:"optional" json:"recipientPropertyKey" yaml:"recipientPropertyKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/share#value DataDatabricksShare#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

