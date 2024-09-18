// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskSqlTaskAlert struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/job#alert_id DataDatabricksJob#alert_id}.
	AlertId *string `field:"required" json:"alertId" yaml:"alertId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/job#pause_subscriptions DataDatabricksJob#pause_subscriptions}.
	PauseSubscriptions interface{} `field:"optional" json:"pauseSubscriptions" yaml:"pauseSubscriptions"`
	// subscriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs/data-sources/job#subscriptions DataDatabricksJob#subscriptions}
	Subscriptions interface{} `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

