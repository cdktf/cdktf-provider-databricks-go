// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskWebhookNotifications struct {
	// on_duration_warning_threshold_exceeded block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/job#on_duration_warning_threshold_exceeded Job#on_duration_warning_threshold_exceeded}
	OnDurationWarningThresholdExceeded interface{} `field:"optional" json:"onDurationWarningThresholdExceeded" yaml:"onDurationWarningThresholdExceeded"`
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/job#on_failure Job#on_failure}
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_start block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/job#on_start Job#on_start}
	OnStart interface{} `field:"optional" json:"onStart" yaml:"onStart"`
	// on_streaming_backlog_exceeded block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/job#on_streaming_backlog_exceeded Job#on_streaming_backlog_exceeded}
	OnStreamingBacklogExceeded interface{} `field:"optional" json:"onStreamingBacklogExceeded" yaml:"onStreamingBacklogExceeded"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/job#on_success Job#on_success}
	OnSuccess interface{} `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

