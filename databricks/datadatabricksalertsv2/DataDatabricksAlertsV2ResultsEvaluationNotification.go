// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertsv2


type DataDatabricksAlertsV2ResultsEvaluationNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/alerts_v2#notify_on_ok DataDatabricksAlertsV2#notify_on_ok}.
	NotifyOnOk interface{} `field:"optional" json:"notifyOnOk" yaml:"notifyOnOk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/alerts_v2#retrigger_seconds DataDatabricksAlertsV2#retrigger_seconds}.
	RetriggerSeconds *float64 `field:"optional" json:"retriggerSeconds" yaml:"retriggerSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/alerts_v2#subscriptions DataDatabricksAlertsV2#subscriptions}.
	Subscriptions interface{} `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

