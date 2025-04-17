// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTriggerTableUpdate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/data-sources/job#table_names DataDatabricksJob#table_names}.
	TableNames *[]*string `field:"required" json:"tableNames" yaml:"tableNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/data-sources/job#condition DataDatabricksJob#condition}.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/data-sources/job#min_time_between_triggers_seconds DataDatabricksJob#min_time_between_triggers_seconds}.
	MinTimeBetweenTriggersSeconds *float64 `field:"optional" json:"minTimeBetweenTriggersSeconds" yaml:"minTimeBetweenTriggersSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/data-sources/job#wait_after_last_change_seconds DataDatabricksJob#wait_after_last_change_seconds}.
	WaitAfterLastChangeSeconds *float64 `field:"optional" json:"waitAfterLastChangeSeconds" yaml:"waitAfterLastChangeSeconds"`
}

