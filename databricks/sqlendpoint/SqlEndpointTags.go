// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlendpoint


type SqlEndpointTags struct {
	// custom_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/sql_endpoint#custom_tags SqlEndpoint#custom_tags}
	CustomTags interface{} `field:"optional" json:"customTags" yaml:"customTags"`
}

