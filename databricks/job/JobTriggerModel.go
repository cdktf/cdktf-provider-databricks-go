// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTriggerModel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#condition Job#condition}.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#aliases Job#aliases}.
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#min_time_between_triggers_seconds Job#min_time_between_triggers_seconds}.
	MinTimeBetweenTriggersSeconds *float64 `field:"optional" json:"minTimeBetweenTriggersSeconds" yaml:"minTimeBetweenTriggersSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#securable_name Job#securable_name}.
	SecurableName *string `field:"optional" json:"securableName" yaml:"securableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/job#wait_after_last_change_seconds Job#wait_after_last_change_seconds}.
	WaitAfterLastChangeSeconds *float64 `field:"optional" json:"waitAfterLastChangeSeconds" yaml:"waitAfterLastChangeSeconds"`
}

