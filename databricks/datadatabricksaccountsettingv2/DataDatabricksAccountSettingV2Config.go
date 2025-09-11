// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountsettingv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksAccountSettingV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#aibi_dashboard_embedding_access_policy DataDatabricksAccountSettingV2#aibi_dashboard_embedding_access_policy}.
	AibiDashboardEmbeddingAccessPolicy *DataDatabricksAccountSettingV2AibiDashboardEmbeddingAccessPolicy `field:"optional" json:"aibiDashboardEmbeddingAccessPolicy" yaml:"aibiDashboardEmbeddingAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#aibi_dashboard_embedding_approved_domains DataDatabricksAccountSettingV2#aibi_dashboard_embedding_approved_domains}.
	AibiDashboardEmbeddingApprovedDomains *DataDatabricksAccountSettingV2AibiDashboardEmbeddingApprovedDomains `field:"optional" json:"aibiDashboardEmbeddingApprovedDomains" yaml:"aibiDashboardEmbeddingApprovedDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#automatic_cluster_update_workspace DataDatabricksAccountSettingV2#automatic_cluster_update_workspace}.
	AutomaticClusterUpdateWorkspace *DataDatabricksAccountSettingV2AutomaticClusterUpdateWorkspace `field:"optional" json:"automaticClusterUpdateWorkspace" yaml:"automaticClusterUpdateWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#boolean_val DataDatabricksAccountSettingV2#boolean_val}.
	BooleanVal *DataDatabricksAccountSettingV2BooleanVal `field:"optional" json:"booleanVal" yaml:"booleanVal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#default_data_security_mode DataDatabricksAccountSettingV2#default_data_security_mode}.
	DefaultDataSecurityMode *DataDatabricksAccountSettingV2DefaultDataSecurityMode `field:"optional" json:"defaultDataSecurityMode" yaml:"defaultDataSecurityMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#effective_aibi_dashboard_embedding_access_policy DataDatabricksAccountSettingV2#effective_aibi_dashboard_embedding_access_policy}.
	EffectiveAibiDashboardEmbeddingAccessPolicy *DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy `field:"optional" json:"effectiveAibiDashboardEmbeddingAccessPolicy" yaml:"effectiveAibiDashboardEmbeddingAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#effective_aibi_dashboard_embedding_approved_domains DataDatabricksAccountSettingV2#effective_aibi_dashboard_embedding_approved_domains}.
	EffectiveAibiDashboardEmbeddingApprovedDomains *DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains `field:"optional" json:"effectiveAibiDashboardEmbeddingApprovedDomains" yaml:"effectiveAibiDashboardEmbeddingApprovedDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#effective_automatic_cluster_update_workspace DataDatabricksAccountSettingV2#effective_automatic_cluster_update_workspace}.
	EffectiveAutomaticClusterUpdateWorkspace *DataDatabricksAccountSettingV2EffectiveAutomaticClusterUpdateWorkspace `field:"optional" json:"effectiveAutomaticClusterUpdateWorkspace" yaml:"effectiveAutomaticClusterUpdateWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#effective_default_data_security_mode DataDatabricksAccountSettingV2#effective_default_data_security_mode}.
	EffectiveDefaultDataSecurityMode *DataDatabricksAccountSettingV2EffectiveDefaultDataSecurityMode `field:"optional" json:"effectiveDefaultDataSecurityMode" yaml:"effectiveDefaultDataSecurityMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#effective_personal_compute DataDatabricksAccountSettingV2#effective_personal_compute}.
	EffectivePersonalCompute *DataDatabricksAccountSettingV2EffectivePersonalCompute `field:"optional" json:"effectivePersonalCompute" yaml:"effectivePersonalCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#effective_restrict_workspace_admins DataDatabricksAccountSettingV2#effective_restrict_workspace_admins}.
	EffectiveRestrictWorkspaceAdmins *DataDatabricksAccountSettingV2EffectiveRestrictWorkspaceAdmins `field:"optional" json:"effectiveRestrictWorkspaceAdmins" yaml:"effectiveRestrictWorkspaceAdmins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#integer_val DataDatabricksAccountSettingV2#integer_val}.
	IntegerVal *DataDatabricksAccountSettingV2IntegerVal `field:"optional" json:"integerVal" yaml:"integerVal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#name DataDatabricksAccountSettingV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#personal_compute DataDatabricksAccountSettingV2#personal_compute}.
	PersonalCompute *DataDatabricksAccountSettingV2PersonalCompute `field:"optional" json:"personalCompute" yaml:"personalCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#restrict_workspace_admins DataDatabricksAccountSettingV2#restrict_workspace_admins}.
	RestrictWorkspaceAdmins *DataDatabricksAccountSettingV2RestrictWorkspaceAdmins `field:"optional" json:"restrictWorkspaceAdmins" yaml:"restrictWorkspaceAdmins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/account_setting_v2#string_val DataDatabricksAccountSettingV2#string_val}.
	StringVal *DataDatabricksAccountSettingV2StringVal `field:"optional" json:"stringVal" yaml:"stringVal"`
}

