// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package permissionassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PermissionAssignmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/permission_assignment#permissions PermissionAssignment#permissions}.
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/permission_assignment#group_name PermissionAssignment#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/permission_assignment#id PermissionAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/permission_assignment#principal_id PermissionAssignment#principal_id}.
	PrincipalId *float64 `field:"optional" json:"principalId" yaml:"principalId"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/permission_assignment#provider_config PermissionAssignment#provider_config}
	ProviderConfig *PermissionAssignmentProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/permission_assignment#service_principal_name PermissionAssignment#service_principal_name}.
	ServicePrincipalName *string `field:"optional" json:"servicePrincipalName" yaml:"servicePrincipalName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/permission_assignment#user_name PermissionAssignment#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

