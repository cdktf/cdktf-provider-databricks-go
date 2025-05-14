// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskDashboardTaskSubscription struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/resources/job#custom_subject Job#custom_subject}.
	CustomSubject *string `field:"optional" json:"customSubject" yaml:"customSubject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/resources/job#paused Job#paused}.
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// subscribers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/resources/job#subscribers Job#subscribers}
	Subscribers interface{} `field:"optional" json:"subscribers" yaml:"subscribers"`
}

