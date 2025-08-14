// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomscleanroom


type CleanRoomsCleanRoomRemoteDetailedInfoCollaborators struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/clean_rooms_clean_room#collaborator_alias CleanRoomsCleanRoom#collaborator_alias}.
	CollaboratorAlias *string `field:"required" json:"collaboratorAlias" yaml:"collaboratorAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/clean_rooms_clean_room#global_metastore_id CleanRoomsCleanRoom#global_metastore_id}.
	GlobalMetastoreId *string `field:"optional" json:"globalMetastoreId" yaml:"globalMetastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/clean_rooms_clean_room#invite_recipient_email CleanRoomsCleanRoom#invite_recipient_email}.
	InviteRecipientEmail *string `field:"optional" json:"inviteRecipientEmail" yaml:"inviteRecipientEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.87.1/docs/resources/clean_rooms_clean_room#invite_recipient_workspace_id CleanRoomsCleanRoom#invite_recipient_workspace_id}.
	InviteRecipientWorkspaceId *float64 `field:"optional" json:"inviteRecipientWorkspaceId" yaml:"inviteRecipientWorkspaceId"`
}

