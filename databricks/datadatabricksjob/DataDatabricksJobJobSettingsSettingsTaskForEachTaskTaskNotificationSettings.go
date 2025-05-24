// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/job#alert_on_last_attempt DataDatabricksJob#alert_on_last_attempt}.
	AlertOnLastAttempt interface{} `field:"optional" json:"alertOnLastAttempt" yaml:"alertOnLastAttempt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/job#no_alert_for_canceled_runs DataDatabricksJob#no_alert_for_canceled_runs}.
	NoAlertForCanceledRuns interface{} `field:"optional" json:"noAlertForCanceledRuns" yaml:"noAlertForCanceledRuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/job#no_alert_for_skipped_runs DataDatabricksJob#no_alert_for_skipped_runs}.
	NoAlertForSkippedRuns interface{} `field:"optional" json:"noAlertForSkippedRuns" yaml:"noAlertForSkippedRuns"`
}

