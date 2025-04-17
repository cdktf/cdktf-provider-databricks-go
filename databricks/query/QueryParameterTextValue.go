// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package query


type QueryParameterTextValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/query#value Query#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

