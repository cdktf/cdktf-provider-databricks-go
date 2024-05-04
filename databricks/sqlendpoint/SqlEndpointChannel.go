// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlendpoint


type SqlEndpointChannel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/sql_endpoint#dbsql_version SqlEndpoint#dbsql_version}.
	DbsqlVersion *string `field:"optional" json:"dbsqlVersion" yaml:"dbsqlVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/sql_endpoint#name SqlEndpoint#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

