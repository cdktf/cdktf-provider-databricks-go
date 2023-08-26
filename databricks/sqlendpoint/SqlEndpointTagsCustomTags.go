// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlendpoint


type SqlEndpointTagsCustomTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.0/docs/resources/sql_endpoint#key SqlEndpoint#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.0/docs/resources/sql_endpoint#value SqlEndpoint#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

