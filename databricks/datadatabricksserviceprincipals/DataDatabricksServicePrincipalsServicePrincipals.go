// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksserviceprincipals


type DataDatabricksServicePrincipalsServicePrincipals struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/service_principals#acl_principal_id DataDatabricksServicePrincipals#acl_principal_id}.
	AclPrincipalId *string `field:"optional" json:"aclPrincipalId" yaml:"aclPrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/service_principals#active DataDatabricksServicePrincipals#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/service_principals#application_id DataDatabricksServicePrincipals#application_id}.
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/service_principals#display_name DataDatabricksServicePrincipals#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/service_principals#external_id DataDatabricksServicePrincipals#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/service_principals#home DataDatabricksServicePrincipals#home}.
	Home *string `field:"optional" json:"home" yaml:"home"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/service_principals#id DataDatabricksServicePrincipals#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/service_principals#repos DataDatabricksServicePrincipals#repos}.
	Repos *string `field:"optional" json:"repos" yaml:"repos"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/service_principals#scim_id DataDatabricksServicePrincipals#scim_id}.
	ScimId *string `field:"optional" json:"scimId" yaml:"scimId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/service_principals#sp_id DataDatabricksServicePrincipals#sp_id}.
	SpId *string `field:"optional" json:"spId" yaml:"spId"`
}

