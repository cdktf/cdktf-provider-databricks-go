// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksusers


type DataDatabricksUsersUsers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/users#active DataDatabricksUsers#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/users#display_name DataDatabricksUsers#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/users#emails DataDatabricksUsers#emails}.
	Emails interface{} `field:"optional" json:"emails" yaml:"emails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/users#entitlements DataDatabricksUsers#entitlements}.
	Entitlements interface{} `field:"optional" json:"entitlements" yaml:"entitlements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/users#external_id DataDatabricksUsers#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/users#groups DataDatabricksUsers#groups}.
	Groups interface{} `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/users#id DataDatabricksUsers#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/users#name DataDatabricksUsers#name}.
	Name *DataDatabricksUsersUsersName `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/users#roles DataDatabricksUsers#roles}.
	Roles interface{} `field:"optional" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/users#schemas DataDatabricksUsers#schemas}.
	Schemas *[]*string `field:"optional" json:"schemas" yaml:"schemas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/users#user_name DataDatabricksUsers#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

