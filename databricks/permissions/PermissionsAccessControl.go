// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package permissions


type PermissionsAccessControl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/resources/permissions#permission_level Permissions#permission_level}.
	PermissionLevel *string `field:"required" json:"permissionLevel" yaml:"permissionLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/resources/permissions#group_name Permissions#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/resources/permissions#service_principal_name Permissions#service_principal_name}.
	ServicePrincipalName *string `field:"optional" json:"servicePrincipalName" yaml:"servicePrincipalName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/resources/permissions#user_name Permissions#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

