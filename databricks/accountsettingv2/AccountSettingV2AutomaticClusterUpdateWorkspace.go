// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2


type AccountSettingV2AutomaticClusterUpdateWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/account_setting_v2#can_toggle AccountSettingV2#can_toggle}.
	CanToggle interface{} `field:"optional" json:"canToggle" yaml:"canToggle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/account_setting_v2#enabled AccountSettingV2#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/account_setting_v2#enablement_details AccountSettingV2#enablement_details}.
	EnablementDetails *AccountSettingV2AutomaticClusterUpdateWorkspaceEnablementDetails `field:"optional" json:"enablementDetails" yaml:"enablementDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/account_setting_v2#maintenance_window AccountSettingV2#maintenance_window}.
	MaintenanceWindow *AccountSettingV2AutomaticClusterUpdateWorkspaceMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/account_setting_v2#restart_even_if_no_updates_available AccountSettingV2#restart_even_if_no_updates_available}.
	RestartEvenIfNoUpdatesAvailable interface{} `field:"optional" json:"restartEvenIfNoUpdatesAvailable" yaml:"restartEvenIfNoUpdatesAvailable"`
}

