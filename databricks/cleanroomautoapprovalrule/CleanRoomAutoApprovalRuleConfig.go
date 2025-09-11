// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomautoapprovalrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanRoomAutoApprovalRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/clean_room_auto_approval_rule#author_collaborator_alias CleanRoomAutoApprovalRule#author_collaborator_alias}.
	AuthorCollaboratorAlias *string `field:"optional" json:"authorCollaboratorAlias" yaml:"authorCollaboratorAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/clean_room_auto_approval_rule#author_scope CleanRoomAutoApprovalRule#author_scope}.
	AuthorScope *string `field:"optional" json:"authorScope" yaml:"authorScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/clean_room_auto_approval_rule#clean_room_name CleanRoomAutoApprovalRule#clean_room_name}.
	CleanRoomName *string `field:"optional" json:"cleanRoomName" yaml:"cleanRoomName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/clean_room_auto_approval_rule#runner_collaborator_alias CleanRoomAutoApprovalRule#runner_collaborator_alias}.
	RunnerCollaboratorAlias *string `field:"optional" json:"runnerCollaboratorAlias" yaml:"runnerCollaboratorAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/clean_room_auto_approval_rule#workspace_id CleanRoomAutoApprovalRule#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

