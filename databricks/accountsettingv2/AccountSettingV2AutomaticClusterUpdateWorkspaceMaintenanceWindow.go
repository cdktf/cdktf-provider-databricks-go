// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2


type AccountSettingV2AutomaticClusterUpdateWorkspaceMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/account_setting_v2#week_day_based_schedule AccountSettingV2#week_day_based_schedule}.
	WeekDayBasedSchedule *AccountSettingV2AutomaticClusterUpdateWorkspaceMaintenanceWindowWeekDayBasedSchedule `field:"optional" json:"weekDayBasedSchedule" yaml:"weekDayBasedSchedule"`
}

