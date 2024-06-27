// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automaticclusterupdateworkspacesetting


type AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/resources/automatic_cluster_update_workspace_setting#can_toggle AutomaticClusterUpdateWorkspaceSetting#can_toggle}.
	CanToggle interface{} `field:"optional" json:"canToggle" yaml:"canToggle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/resources/automatic_cluster_update_workspace_setting#enabled AutomaticClusterUpdateWorkspaceSetting#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// enablement_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/resources/automatic_cluster_update_workspace_setting#enablement_details AutomaticClusterUpdateWorkspaceSetting#enablement_details}
	EnablementDetails *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetails `field:"optional" json:"enablementDetails" yaml:"enablementDetails"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/resources/automatic_cluster_update_workspace_setting#maintenance_window AutomaticClusterUpdateWorkspaceSetting#maintenance_window}
	MaintenanceWindow *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/resources/automatic_cluster_update_workspace_setting#restart_even_if_no_updates_available AutomaticClusterUpdateWorkspaceSetting#restart_even_if_no_updates_available}.
	RestartEvenIfNoUpdatesAvailable interface{} `field:"optional" json:"restartEvenIfNoUpdatesAvailable" yaml:"restartEvenIfNoUpdatesAvailable"`
}

