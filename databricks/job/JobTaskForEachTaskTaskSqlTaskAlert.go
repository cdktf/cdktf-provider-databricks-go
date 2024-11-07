// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskSqlTaskAlert struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/job#alert_id Job#alert_id}.
	AlertId *string `field:"required" json:"alertId" yaml:"alertId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/job#pause_subscriptions Job#pause_subscriptions}.
	PauseSubscriptions interface{} `field:"optional" json:"pauseSubscriptions" yaml:"pauseSubscriptions"`
	// subscriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/job#subscriptions Job#subscriptions}
	Subscriptions interface{} `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

