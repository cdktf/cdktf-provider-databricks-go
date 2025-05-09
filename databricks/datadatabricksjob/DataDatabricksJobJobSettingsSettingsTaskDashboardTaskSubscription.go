// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskDashboardTaskSubscription struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/job#custom_subject DataDatabricksJob#custom_subject}.
	CustomSubject *string `field:"optional" json:"customSubject" yaml:"customSubject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/job#paused DataDatabricksJob#paused}.
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// subscribers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/job#subscribers DataDatabricksJob#subscribers}
	Subscribers interface{} `field:"optional" json:"subscribers" yaml:"subscribers"`
}

