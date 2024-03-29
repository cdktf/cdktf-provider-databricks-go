// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlendpoint


type SqlEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/sql_endpoint#create SqlEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

