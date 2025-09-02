// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomasset


type CleanRoomAssetNotebookReviews struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#comment CleanRoomAsset#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#created_at_millis CleanRoomAsset#created_at_millis}.
	CreatedAtMillis *float64 `field:"optional" json:"createdAtMillis" yaml:"createdAtMillis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#reviewer_collaborator_alias CleanRoomAsset#reviewer_collaborator_alias}.
	ReviewerCollaboratorAlias *string `field:"optional" json:"reviewerCollaboratorAlias" yaml:"reviewerCollaboratorAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#review_state CleanRoomAsset#review_state}.
	ReviewState *string `field:"optional" json:"reviewState" yaml:"reviewState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/clean_room_asset#review_sub_reason CleanRoomAsset#review_sub_reason}.
	ReviewSubReason *string `field:"optional" json:"reviewSubReason" yaml:"reviewSubReason"`
}

