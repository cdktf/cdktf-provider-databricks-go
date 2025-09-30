// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspacesettingv2


type DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/workspace_setting_v2#can_toggle DataDatabricksWorkspaceSettingV2#can_toggle}.
	CanToggle interface{} `field:"optional" json:"canToggle" yaml:"canToggle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/workspace_setting_v2#enabled DataDatabricksWorkspaceSettingV2#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/workspace_setting_v2#enablement_details DataDatabricksWorkspaceSettingV2#enablement_details}.
	EnablementDetails *DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetails `field:"optional" json:"enablementDetails" yaml:"enablementDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/workspace_setting_v2#maintenance_window DataDatabricksWorkspaceSettingV2#maintenance_window}.
	MaintenanceWindow *DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspaceMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/workspace_setting_v2#restart_even_if_no_updates_available DataDatabricksWorkspaceSettingV2#restart_even_if_no_updates_available}.
	RestartEvenIfNoUpdatesAvailable interface{} `field:"optional" json:"restartEvenIfNoUpdatesAvailable" yaml:"restartEvenIfNoUpdatesAvailable"`
}

