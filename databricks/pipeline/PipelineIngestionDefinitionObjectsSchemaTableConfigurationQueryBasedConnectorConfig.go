// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaTableConfigurationQueryBasedConnectorConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#cursor_columns Pipeline#cursor_columns}.
	CursorColumns *[]*string `field:"optional" json:"cursorColumns" yaml:"cursorColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#deletion_condition Pipeline#deletion_condition}.
	DeletionCondition *string `field:"optional" json:"deletionCondition" yaml:"deletionCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#hard_deletion_sync_min_interval_in_seconds Pipeline#hard_deletion_sync_min_interval_in_seconds}.
	HardDeletionSyncMinIntervalInSeconds *float64 `field:"optional" json:"hardDeletionSyncMinIntervalInSeconds" yaml:"hardDeletionSyncMinIntervalInSeconds"`
}

