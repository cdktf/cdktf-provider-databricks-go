// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automaticclusterupdateworkspacesetting


type AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/resources/automatic_cluster_update_workspace_setting#forced_for_compliance_mode AutomaticClusterUpdateWorkspaceSetting#forced_for_compliance_mode}.
	ForcedForComplianceMode interface{} `field:"optional" json:"forcedForComplianceMode" yaml:"forcedForComplianceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/resources/automatic_cluster_update_workspace_setting#unavailable_for_disabled_entitlement AutomaticClusterUpdateWorkspaceSetting#unavailable_for_disabled_entitlement}.
	UnavailableForDisabledEntitlement interface{} `field:"optional" json:"unavailableForDisabledEntitlement" yaml:"unavailableForDisabledEntitlement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/resources/automatic_cluster_update_workspace_setting#unavailable_for_non_enterprise_tier AutomaticClusterUpdateWorkspaceSetting#unavailable_for_non_enterprise_tier}.
	UnavailableForNonEnterpriseTier interface{} `field:"optional" json:"unavailableForNonEnterpriseTier" yaml:"unavailableForNonEnterpriseTier"`
}

