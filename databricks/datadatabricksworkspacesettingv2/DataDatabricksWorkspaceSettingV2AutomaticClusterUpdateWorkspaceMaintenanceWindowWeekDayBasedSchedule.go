// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspacesettingv2


type DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspaceMaintenanceWindowWeekDayBasedSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/workspace_setting_v2#day_of_week DataDatabricksWorkspaceSettingV2#day_of_week}.
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/workspace_setting_v2#frequency DataDatabricksWorkspaceSettingV2#frequency}.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/workspace_setting_v2#window_start_time DataDatabricksWorkspaceSettingV2#window_start_time}.
	WindowStartTime *DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspaceMaintenanceWindowWeekDayBasedScheduleWindowStartTime `field:"optional" json:"windowStartTime" yaml:"windowStartTime"`
}

