// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automaticclusterupdateworkspacesetting


type AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindowWeekDayBasedScheduleWindowStartTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/resources/automatic_cluster_update_workspace_setting#hours AutomaticClusterUpdateWorkspaceSetting#hours}.
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/resources/automatic_cluster_update_workspace_setting#minutes AutomaticClusterUpdateWorkspaceSetting#minutes}.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
}
