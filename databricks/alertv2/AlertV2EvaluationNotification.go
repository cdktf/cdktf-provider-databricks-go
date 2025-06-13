// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertv2


type AlertV2EvaluationNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/alert_v2#notify_on_ok AlertV2#notify_on_ok}.
	NotifyOnOk interface{} `field:"optional" json:"notifyOnOk" yaml:"notifyOnOk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/alert_v2#retrigger_seconds AlertV2#retrigger_seconds}.
	RetriggerSeconds *float64 `field:"optional" json:"retriggerSeconds" yaml:"retriggerSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/alert_v2#subscriptions AlertV2#subscriptions}.
	Subscriptions interface{} `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

