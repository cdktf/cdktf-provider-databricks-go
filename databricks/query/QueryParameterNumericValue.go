// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query


type QueryParameterNumericValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/query#value Query#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

