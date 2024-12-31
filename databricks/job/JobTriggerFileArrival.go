// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTriggerFileArrival struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/job#url Job#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/job#min_time_between_triggers_seconds Job#min_time_between_triggers_seconds}.
	MinTimeBetweenTriggersSeconds *float64 `field:"optional" json:"minTimeBetweenTriggersSeconds" yaml:"minTimeBetweenTriggersSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/resources/job#wait_after_last_change_seconds Job#wait_after_last_change_seconds}.
	WaitAfterLastChangeSeconds *float64 `field:"optional" json:"waitAfterLastChangeSeconds" yaml:"waitAfterLastChangeSeconds"`
}

