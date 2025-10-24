// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountsettingv2


type DataDatabricksAccountSettingV2AutomaticClusterUpdateWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/account_setting_v2#can_toggle DataDatabricksAccountSettingV2#can_toggle}.
	CanToggle interface{} `field:"optional" json:"canToggle" yaml:"canToggle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/account_setting_v2#enabled DataDatabricksAccountSettingV2#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/account_setting_v2#enablement_details DataDatabricksAccountSettingV2#enablement_details}.
	EnablementDetails *DataDatabricksAccountSettingV2AutomaticClusterUpdateWorkspaceEnablementDetails `field:"optional" json:"enablementDetails" yaml:"enablementDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/account_setting_v2#maintenance_window DataDatabricksAccountSettingV2#maintenance_window}.
	MaintenanceWindow *DataDatabricksAccountSettingV2AutomaticClusterUpdateWorkspaceMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/account_setting_v2#restart_even_if_no_updates_available DataDatabricksAccountSettingV2#restart_even_if_no_updates_available}.
	RestartEvenIfNoUpdatesAvailable interface{} `field:"optional" json:"restartEvenIfNoUpdatesAvailable" yaml:"restartEvenIfNoUpdatesAvailable"`
}

