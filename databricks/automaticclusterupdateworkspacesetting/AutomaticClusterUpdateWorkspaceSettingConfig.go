// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automaticclusterupdateworkspacesetting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomaticClusterUpdateWorkspaceSettingConfig struct {
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
	// automatic_cluster_update_workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/automatic_cluster_update_workspace_setting#automatic_cluster_update_workspace AutomaticClusterUpdateWorkspaceSetting#automatic_cluster_update_workspace}
	AutomaticClusterUpdateWorkspace *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspace `field:"required" json:"automaticClusterUpdateWorkspace" yaml:"automaticClusterUpdateWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/automatic_cluster_update_workspace_setting#etag AutomaticClusterUpdateWorkspaceSetting#etag}.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/automatic_cluster_update_workspace_setting#id AutomaticClusterUpdateWorkspaceSetting#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/automatic_cluster_update_workspace_setting#setting_name AutomaticClusterUpdateWorkspaceSetting#setting_name}.
	SettingName *string `field:"optional" json:"settingName" yaml:"settingName"`
}

