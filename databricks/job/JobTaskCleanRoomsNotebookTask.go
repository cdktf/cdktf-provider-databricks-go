// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskCleanRoomsNotebookTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/job#clean_room_name Job#clean_room_name}.
	CleanRoomName *string `field:"required" json:"cleanRoomName" yaml:"cleanRoomName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/job#notebook_name Job#notebook_name}.
	NotebookName *string `field:"required" json:"notebookName" yaml:"notebookName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/job#etag Job#etag}.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/job#notebook_base_parameters Job#notebook_base_parameters}.
	NotebookBaseParameters *map[string]*string `field:"optional" json:"notebookBaseParameters" yaml:"notebookBaseParameters"`
}

