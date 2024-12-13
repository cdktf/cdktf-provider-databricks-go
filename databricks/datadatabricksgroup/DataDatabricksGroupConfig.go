// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#display_name DataDatabricksGroup#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#acl_principal_id DataDatabricksGroup#acl_principal_id}.
	AclPrincipalId *string `field:"optional" json:"aclPrincipalId" yaml:"aclPrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#allow_cluster_create DataDatabricksGroup#allow_cluster_create}.
	AllowClusterCreate interface{} `field:"optional" json:"allowClusterCreate" yaml:"allowClusterCreate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#allow_instance_pool_create DataDatabricksGroup#allow_instance_pool_create}.
	AllowInstancePoolCreate interface{} `field:"optional" json:"allowInstancePoolCreate" yaml:"allowInstancePoolCreate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#child_groups DataDatabricksGroup#child_groups}.
	ChildGroups *[]*string `field:"optional" json:"childGroups" yaml:"childGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#databricks_sql_access DataDatabricksGroup#databricks_sql_access}.
	DatabricksSqlAccess interface{} `field:"optional" json:"databricksSqlAccess" yaml:"databricksSqlAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#external_id DataDatabricksGroup#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#groups DataDatabricksGroup#groups}.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#id DataDatabricksGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#instance_profiles DataDatabricksGroup#instance_profiles}.
	InstanceProfiles *[]*string `field:"optional" json:"instanceProfiles" yaml:"instanceProfiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#members DataDatabricksGroup#members}.
	Members *[]*string `field:"optional" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#recursive DataDatabricksGroup#recursive}.
	Recursive interface{} `field:"optional" json:"recursive" yaml:"recursive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#service_principals DataDatabricksGroup#service_principals}.
	ServicePrincipals *[]*string `field:"optional" json:"servicePrincipals" yaml:"servicePrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#users DataDatabricksGroup#users}.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/group#workspace_access DataDatabricksGroup#workspace_access}.
	WorkspaceAccess interface{} `field:"optional" json:"workspaceAccess" yaml:"workspaceAccess"`
}

