// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package group

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/group#display_name Group#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/group#acl_principal_id Group#acl_principal_id}.
	AclPrincipalId *string `field:"optional" json:"aclPrincipalId" yaml:"aclPrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/group#allow_cluster_create Group#allow_cluster_create}.
	AllowClusterCreate interface{} `field:"optional" json:"allowClusterCreate" yaml:"allowClusterCreate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/group#allow_instance_pool_create Group#allow_instance_pool_create}.
	AllowInstancePoolCreate interface{} `field:"optional" json:"allowInstancePoolCreate" yaml:"allowInstancePoolCreate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/group#databricks_sql_access Group#databricks_sql_access}.
	DatabricksSqlAccess interface{} `field:"optional" json:"databricksSqlAccess" yaml:"databricksSqlAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/group#external_id Group#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/group#force Group#force}.
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/group#id Group#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/group#url Group#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/resources/group#workspace_access Group#workspace_access}.
	WorkspaceAccess interface{} `field:"optional" json:"workspaceAccess" yaml:"workspaceAccess"`
}

