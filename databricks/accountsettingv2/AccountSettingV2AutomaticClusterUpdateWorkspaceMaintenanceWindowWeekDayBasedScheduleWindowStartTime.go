// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2


type AccountSettingV2AutomaticClusterUpdateWorkspaceMaintenanceWindowWeekDayBasedScheduleWindowStartTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/account_setting_v2#hours AccountSettingV2#hours}.
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/account_setting_v2#minutes AccountSettingV2#minutes}.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
}

