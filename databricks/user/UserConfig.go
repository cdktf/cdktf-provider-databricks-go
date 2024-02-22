// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#user_name User#user_name}.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#acl_principal_id User#acl_principal_id}.
	AclPrincipalId *string `field:"optional" json:"aclPrincipalId" yaml:"aclPrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#active User#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#allow_cluster_create User#allow_cluster_create}.
	AllowClusterCreate interface{} `field:"optional" json:"allowClusterCreate" yaml:"allowClusterCreate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#allow_instance_pool_create User#allow_instance_pool_create}.
	AllowInstancePoolCreate interface{} `field:"optional" json:"allowInstancePoolCreate" yaml:"allowInstancePoolCreate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#databricks_sql_access User#databricks_sql_access}.
	DatabricksSqlAccess interface{} `field:"optional" json:"databricksSqlAccess" yaml:"databricksSqlAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#disable_as_user_deletion User#disable_as_user_deletion}.
	DisableAsUserDeletion interface{} `field:"optional" json:"disableAsUserDeletion" yaml:"disableAsUserDeletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#display_name User#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#external_id User#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#force User#force}.
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#force_delete_home_dir User#force_delete_home_dir}.
	ForceDeleteHomeDir interface{} `field:"optional" json:"forceDeleteHomeDir" yaml:"forceDeleteHomeDir"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#force_delete_repos User#force_delete_repos}.
	ForceDeleteRepos interface{} `field:"optional" json:"forceDeleteRepos" yaml:"forceDeleteRepos"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#home User#home}.
	Home *string `field:"optional" json:"home" yaml:"home"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#id User#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#repos User#repos}.
	Repos *string `field:"optional" json:"repos" yaml:"repos"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/resources/user#workspace_access User#workspace_access}.
	WorkspaceAccess interface{} `field:"optional" json:"workspaceAccess" yaml:"workspaceAccess"`
}

