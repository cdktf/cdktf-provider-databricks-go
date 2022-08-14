// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type PermissionsAccessControl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#permission_level Permissions#permission_level}.
	PermissionLevel *string `field:"required" json:"permissionLevel" yaml:"permissionLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#group_name Permissions#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#service_principal_name Permissions#service_principal_name}.
	ServicePrincipalName *string `field:"optional" json:"servicePrincipalName" yaml:"servicePrincipalName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#user_name Permissions#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

