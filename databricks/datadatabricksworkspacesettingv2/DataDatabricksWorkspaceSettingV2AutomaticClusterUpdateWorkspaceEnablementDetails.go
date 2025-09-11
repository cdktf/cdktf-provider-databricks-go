// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspacesettingv2


type DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/workspace_setting_v2#forced_for_compliance_mode DataDatabricksWorkspaceSettingV2#forced_for_compliance_mode}.
	ForcedForComplianceMode interface{} `field:"optional" json:"forcedForComplianceMode" yaml:"forcedForComplianceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/workspace_setting_v2#unavailable_for_disabled_entitlement DataDatabricksWorkspaceSettingV2#unavailable_for_disabled_entitlement}.
	UnavailableForDisabledEntitlement interface{} `field:"optional" json:"unavailableForDisabledEntitlement" yaml:"unavailableForDisabledEntitlement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/workspace_setting_v2#unavailable_for_non_enterprise_tier DataDatabricksWorkspaceSettingV2#unavailable_for_non_enterprise_tier}.
	UnavailableForNonEnterpriseTier interface{} `field:"optional" json:"unavailableForNonEnterpriseTier" yaml:"unavailableForNonEnterpriseTier"`
}

