// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package catalog


type CatalogEffectivePredictiveOptimizationFlag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/catalog#value Catalog#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/catalog#inherited_from_name Catalog#inherited_from_name}.
	InheritedFromName *string `field:"optional" json:"inheritedFromName" yaml:"inheritedFromName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/catalog#inherited_from_type Catalog#inherited_from_type}.
	InheritedFromType *string `field:"optional" json:"inheritedFromType" yaml:"inheritedFromType"`
}

