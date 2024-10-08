// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package share


type ShareObjectPartitionValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/share#name Share#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/share#op Share#op}.
	Op *string `field:"required" json:"op" yaml:"op"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/share#recipient_property_key Share#recipient_property_key}.
	RecipientPropertyKey *string `field:"optional" json:"recipientPropertyKey" yaml:"recipientPropertyKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/share#value Share#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

