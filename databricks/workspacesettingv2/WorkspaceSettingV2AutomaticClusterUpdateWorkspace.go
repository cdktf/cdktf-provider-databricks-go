// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2AutomaticClusterUpdateWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/workspace_setting_v2#can_toggle WorkspaceSettingV2#can_toggle}.
	CanToggle interface{} `field:"optional" json:"canToggle" yaml:"canToggle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/workspace_setting_v2#enabled WorkspaceSettingV2#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/workspace_setting_v2#enablement_details WorkspaceSettingV2#enablement_details}.
	EnablementDetails *WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetails `field:"optional" json:"enablementDetails" yaml:"enablementDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/workspace_setting_v2#maintenance_window WorkspaceSettingV2#maintenance_window}.
	MaintenanceWindow *WorkspaceSettingV2AutomaticClusterUpdateWorkspaceMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/workspace_setting_v2#restart_even_if_no_updates_available WorkspaceSettingV2#restart_even_if_no_updates_available}.
	RestartEvenIfNoUpdatesAvailable interface{} `field:"optional" json:"restartEvenIfNoUpdatesAvailable" yaml:"restartEvenIfNoUpdatesAvailable"`
}

