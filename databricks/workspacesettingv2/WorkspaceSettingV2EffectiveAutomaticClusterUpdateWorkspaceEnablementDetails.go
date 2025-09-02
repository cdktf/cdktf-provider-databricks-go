// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceEnablementDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/workspace_setting_v2#forced_for_compliance_mode WorkspaceSettingV2#forced_for_compliance_mode}.
	ForcedForComplianceMode interface{} `field:"optional" json:"forcedForComplianceMode" yaml:"forcedForComplianceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/workspace_setting_v2#unavailable_for_disabled_entitlement WorkspaceSettingV2#unavailable_for_disabled_entitlement}.
	UnavailableForDisabledEntitlement interface{} `field:"optional" json:"unavailableForDisabledEntitlement" yaml:"unavailableForDisabledEntitlement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/resources/workspace_setting_v2#unavailable_for_non_enterprise_tier WorkspaceSettingV2#unavailable_for_non_enterprise_tier}.
	UnavailableForNonEnterpriseTier interface{} `field:"optional" json:"unavailableForNonEnterpriseTier" yaml:"unavailableForNonEnterpriseTier"`
}

