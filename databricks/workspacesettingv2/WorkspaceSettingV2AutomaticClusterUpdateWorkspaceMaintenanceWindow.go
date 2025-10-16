// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2AutomaticClusterUpdateWorkspaceMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/workspace_setting_v2#week_day_based_schedule WorkspaceSettingV2#week_day_based_schedule}.
	WeekDayBasedSchedule *WorkspaceSettingV2AutomaticClusterUpdateWorkspaceMaintenanceWindowWeekDayBasedSchedule `field:"optional" json:"weekDayBasedSchedule" yaml:"weekDayBasedSchedule"`
}

