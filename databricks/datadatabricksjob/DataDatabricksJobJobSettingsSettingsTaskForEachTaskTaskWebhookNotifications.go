// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskWebhookNotifications struct {
	// on_duration_warning_threshold_exceeded block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.38.0/docs/data-sources/job#on_duration_warning_threshold_exceeded DataDatabricksJob#on_duration_warning_threshold_exceeded}
	OnDurationWarningThresholdExceeded interface{} `field:"optional" json:"onDurationWarningThresholdExceeded" yaml:"onDurationWarningThresholdExceeded"`
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.38.0/docs/data-sources/job#on_failure DataDatabricksJob#on_failure}
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_start block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.38.0/docs/data-sources/job#on_start DataDatabricksJob#on_start}
	OnStart interface{} `field:"optional" json:"onStart" yaml:"onStart"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.38.0/docs/data-sources/job#on_success DataDatabricksJob#on_success}
	OnSuccess interface{} `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

