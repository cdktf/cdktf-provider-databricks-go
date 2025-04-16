// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobNotificationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/job#no_alert_for_canceled_runs Job#no_alert_for_canceled_runs}.
	NoAlertForCanceledRuns interface{} `field:"optional" json:"noAlertForCanceledRuns" yaml:"noAlertForCanceledRuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/job#no_alert_for_skipped_runs Job#no_alert_for_skipped_runs}.
	NoAlertForSkippedRuns interface{} `field:"optional" json:"noAlertForSkippedRuns" yaml:"noAlertForSkippedRuns"`
}

