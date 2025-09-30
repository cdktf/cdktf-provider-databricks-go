// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountSettingV2Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#aibi_dashboard_embedding_access_policy AccountSettingV2#aibi_dashboard_embedding_access_policy}.
	AibiDashboardEmbeddingAccessPolicy *AccountSettingV2AibiDashboardEmbeddingAccessPolicy `field:"optional" json:"aibiDashboardEmbeddingAccessPolicy" yaml:"aibiDashboardEmbeddingAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#aibi_dashboard_embedding_approved_domains AccountSettingV2#aibi_dashboard_embedding_approved_domains}.
	AibiDashboardEmbeddingApprovedDomains *AccountSettingV2AibiDashboardEmbeddingApprovedDomains `field:"optional" json:"aibiDashboardEmbeddingApprovedDomains" yaml:"aibiDashboardEmbeddingApprovedDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#automatic_cluster_update_workspace AccountSettingV2#automatic_cluster_update_workspace}.
	AutomaticClusterUpdateWorkspace *AccountSettingV2AutomaticClusterUpdateWorkspace `field:"optional" json:"automaticClusterUpdateWorkspace" yaml:"automaticClusterUpdateWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#boolean_val AccountSettingV2#boolean_val}.
	BooleanVal *AccountSettingV2BooleanVal `field:"optional" json:"booleanVal" yaml:"booleanVal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#effective_aibi_dashboard_embedding_access_policy AccountSettingV2#effective_aibi_dashboard_embedding_access_policy}.
	EffectiveAibiDashboardEmbeddingAccessPolicy *AccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy `field:"optional" json:"effectiveAibiDashboardEmbeddingAccessPolicy" yaml:"effectiveAibiDashboardEmbeddingAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#effective_aibi_dashboard_embedding_approved_domains AccountSettingV2#effective_aibi_dashboard_embedding_approved_domains}.
	EffectiveAibiDashboardEmbeddingApprovedDomains *AccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains `field:"optional" json:"effectiveAibiDashboardEmbeddingApprovedDomains" yaml:"effectiveAibiDashboardEmbeddingApprovedDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#effective_automatic_cluster_update_workspace AccountSettingV2#effective_automatic_cluster_update_workspace}.
	EffectiveAutomaticClusterUpdateWorkspace *AccountSettingV2EffectiveAutomaticClusterUpdateWorkspace `field:"optional" json:"effectiveAutomaticClusterUpdateWorkspace" yaml:"effectiveAutomaticClusterUpdateWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#effective_personal_compute AccountSettingV2#effective_personal_compute}.
	EffectivePersonalCompute *AccountSettingV2EffectivePersonalCompute `field:"optional" json:"effectivePersonalCompute" yaml:"effectivePersonalCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#effective_restrict_workspace_admins AccountSettingV2#effective_restrict_workspace_admins}.
	EffectiveRestrictWorkspaceAdmins *AccountSettingV2EffectiveRestrictWorkspaceAdmins `field:"optional" json:"effectiveRestrictWorkspaceAdmins" yaml:"effectiveRestrictWorkspaceAdmins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#integer_val AccountSettingV2#integer_val}.
	IntegerVal *AccountSettingV2IntegerVal `field:"optional" json:"integerVal" yaml:"integerVal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#name AccountSettingV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#personal_compute AccountSettingV2#personal_compute}.
	PersonalCompute *AccountSettingV2PersonalCompute `field:"optional" json:"personalCompute" yaml:"personalCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#restrict_workspace_admins AccountSettingV2#restrict_workspace_admins}.
	RestrictWorkspaceAdmins *AccountSettingV2RestrictWorkspaceAdmins `field:"optional" json:"restrictWorkspaceAdmins" yaml:"restrictWorkspaceAdmins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/account_setting_v2#string_val AccountSettingV2#string_val}.
	StringVal *AccountSettingV2StringVal `field:"optional" json:"stringVal" yaml:"stringVal"`
}

