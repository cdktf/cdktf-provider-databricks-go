// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsEmailNotifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/data-sources/job#no_alert_for_skipped_runs DataDatabricksJob#no_alert_for_skipped_runs}.
	NoAlertForSkippedRuns interface{} `field:"optional" json:"noAlertForSkippedRuns" yaml:"noAlertForSkippedRuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/data-sources/job#on_duration_warning_threshold_exceeded DataDatabricksJob#on_duration_warning_threshold_exceeded}.
	OnDurationWarningThresholdExceeded *[]*string `field:"optional" json:"onDurationWarningThresholdExceeded" yaml:"onDurationWarningThresholdExceeded"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/data-sources/job#on_failure DataDatabricksJob#on_failure}.
	OnFailure *[]*string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/data-sources/job#on_start DataDatabricksJob#on_start}.
	OnStart *[]*string `field:"optional" json:"onStart" yaml:"onStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/data-sources/job#on_streaming_backlog_exceeded DataDatabricksJob#on_streaming_backlog_exceeded}.
	OnStreamingBacklogExceeded *[]*string `field:"optional" json:"onStreamingBacklogExceeded" yaml:"onStreamingBacklogExceeded"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/data-sources/job#on_success DataDatabricksJob#on_success}.
	OnSuccess *[]*string `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

