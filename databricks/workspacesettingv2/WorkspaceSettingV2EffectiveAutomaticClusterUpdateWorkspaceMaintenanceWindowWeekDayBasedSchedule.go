// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceMaintenanceWindowWeekDayBasedSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/workspace_setting_v2#day_of_week WorkspaceSettingV2#day_of_week}.
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/workspace_setting_v2#frequency WorkspaceSettingV2#frequency}.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/workspace_setting_v2#window_start_time WorkspaceSettingV2#window_start_time}.
	WindowStartTime *WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceMaintenanceWindowWeekDayBasedScheduleWindowStartTime `field:"optional" json:"windowStartTime" yaml:"windowStartTime"`
}

