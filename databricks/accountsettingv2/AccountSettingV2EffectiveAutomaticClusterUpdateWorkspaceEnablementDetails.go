// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2


type AccountSettingV2EffectiveAutomaticClusterUpdateWorkspaceEnablementDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/account_setting_v2#forced_for_compliance_mode AccountSettingV2#forced_for_compliance_mode}.
	ForcedForComplianceMode interface{} `field:"optional" json:"forcedForComplianceMode" yaml:"forcedForComplianceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/account_setting_v2#unavailable_for_disabled_entitlement AccountSettingV2#unavailable_for_disabled_entitlement}.
	UnavailableForDisabledEntitlement interface{} `field:"optional" json:"unavailableForDisabledEntitlement" yaml:"unavailableForDisabledEntitlement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/account_setting_v2#unavailable_for_non_enterprise_tier AccountSettingV2#unavailable_for_non_enterprise_tier}.
	UnavailableForNonEnterpriseTier interface{} `field:"optional" json:"unavailableForNonEnterpriseTier" yaml:"unavailableForNonEnterpriseTier"`
}

