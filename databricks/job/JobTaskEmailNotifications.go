// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskEmailNotifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/job#no_alert_for_skipped_runs Job#no_alert_for_skipped_runs}.
	NoAlertForSkippedRuns interface{} `field:"optional" json:"noAlertForSkippedRuns" yaml:"noAlertForSkippedRuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/job#on_duration_warning_threshold_exceeded Job#on_duration_warning_threshold_exceeded}.
	OnDurationWarningThresholdExceeded *[]*string `field:"optional" json:"onDurationWarningThresholdExceeded" yaml:"onDurationWarningThresholdExceeded"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/job#on_failure Job#on_failure}.
	OnFailure *[]*string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/job#on_start Job#on_start}.
	OnStart *[]*string `field:"optional" json:"onStart" yaml:"onStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/job#on_streaming_backlog_exceeded Job#on_streaming_backlog_exceeded}.
	OnStreamingBacklogExceeded *[]*string `field:"optional" json:"onStreamingBacklogExceeded" yaml:"onStreamingBacklogExceeded"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/resources/job#on_success Job#on_success}.
	OnSuccess *[]*string `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

