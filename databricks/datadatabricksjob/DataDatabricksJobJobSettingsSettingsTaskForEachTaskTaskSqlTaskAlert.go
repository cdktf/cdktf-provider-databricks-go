// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskSqlTaskAlert struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs/data-sources/job#alert_id DataDatabricksJob#alert_id}.
	AlertId *string `field:"required" json:"alertId" yaml:"alertId"`
	// subscriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs/data-sources/job#subscriptions DataDatabricksJob#subscriptions}
	Subscriptions interface{} `field:"required" json:"subscriptions" yaml:"subscriptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs/data-sources/job#pause_subscriptions DataDatabricksJob#pause_subscriptions}.
	PauseSubscriptions interface{} `field:"optional" json:"pauseSubscriptions" yaml:"pauseSubscriptions"`
}

