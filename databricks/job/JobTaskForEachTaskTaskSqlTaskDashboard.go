// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskSqlTaskDashboard struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/resources/job#dashboard_id Job#dashboard_id}.
	DashboardId *string `field:"required" json:"dashboardId" yaml:"dashboardId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/resources/job#custom_subject Job#custom_subject}.
	CustomSubject *string `field:"optional" json:"customSubject" yaml:"customSubject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/resources/job#pause_subscriptions Job#pause_subscriptions}.
	PauseSubscriptions interface{} `field:"optional" json:"pauseSubscriptions" yaml:"pauseSubscriptions"`
	// subscriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/resources/job#subscriptions Job#subscriptions}
	Subscriptions interface{} `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

