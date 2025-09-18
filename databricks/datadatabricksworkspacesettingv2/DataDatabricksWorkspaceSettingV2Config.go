// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspacesettingv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksWorkspaceSettingV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#aibi_dashboard_embedding_access_policy DataDatabricksWorkspaceSettingV2#aibi_dashboard_embedding_access_policy}.
	AibiDashboardEmbeddingAccessPolicy *DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingAccessPolicy `field:"optional" json:"aibiDashboardEmbeddingAccessPolicy" yaml:"aibiDashboardEmbeddingAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#aibi_dashboard_embedding_approved_domains DataDatabricksWorkspaceSettingV2#aibi_dashboard_embedding_approved_domains}.
	AibiDashboardEmbeddingApprovedDomains *DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingApprovedDomains `field:"optional" json:"aibiDashboardEmbeddingApprovedDomains" yaml:"aibiDashboardEmbeddingApprovedDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#automatic_cluster_update_workspace DataDatabricksWorkspaceSettingV2#automatic_cluster_update_workspace}.
	AutomaticClusterUpdateWorkspace *DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspace `field:"optional" json:"automaticClusterUpdateWorkspace" yaml:"automaticClusterUpdateWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#boolean_val DataDatabricksWorkspaceSettingV2#boolean_val}.
	BooleanVal *DataDatabricksWorkspaceSettingV2BooleanVal `field:"optional" json:"booleanVal" yaml:"booleanVal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#default_data_security_mode DataDatabricksWorkspaceSettingV2#default_data_security_mode}.
	DefaultDataSecurityMode *DataDatabricksWorkspaceSettingV2DefaultDataSecurityMode `field:"optional" json:"defaultDataSecurityMode" yaml:"defaultDataSecurityMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#effective_aibi_dashboard_embedding_access_policy DataDatabricksWorkspaceSettingV2#effective_aibi_dashboard_embedding_access_policy}.
	EffectiveAibiDashboardEmbeddingAccessPolicy *DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy `field:"optional" json:"effectiveAibiDashboardEmbeddingAccessPolicy" yaml:"effectiveAibiDashboardEmbeddingAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#effective_aibi_dashboard_embedding_approved_domains DataDatabricksWorkspaceSettingV2#effective_aibi_dashboard_embedding_approved_domains}.
	EffectiveAibiDashboardEmbeddingApprovedDomains *DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains `field:"optional" json:"effectiveAibiDashboardEmbeddingApprovedDomains" yaml:"effectiveAibiDashboardEmbeddingApprovedDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#effective_automatic_cluster_update_workspace DataDatabricksWorkspaceSettingV2#effective_automatic_cluster_update_workspace}.
	EffectiveAutomaticClusterUpdateWorkspace *DataDatabricksWorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspace `field:"optional" json:"effectiveAutomaticClusterUpdateWorkspace" yaml:"effectiveAutomaticClusterUpdateWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#effective_default_data_security_mode DataDatabricksWorkspaceSettingV2#effective_default_data_security_mode}.
	EffectiveDefaultDataSecurityMode *DataDatabricksWorkspaceSettingV2EffectiveDefaultDataSecurityMode `field:"optional" json:"effectiveDefaultDataSecurityMode" yaml:"effectiveDefaultDataSecurityMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#effective_personal_compute DataDatabricksWorkspaceSettingV2#effective_personal_compute}.
	EffectivePersonalCompute *DataDatabricksWorkspaceSettingV2EffectivePersonalCompute `field:"optional" json:"effectivePersonalCompute" yaml:"effectivePersonalCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#effective_restrict_workspace_admins DataDatabricksWorkspaceSettingV2#effective_restrict_workspace_admins}.
	EffectiveRestrictWorkspaceAdmins *DataDatabricksWorkspaceSettingV2EffectiveRestrictWorkspaceAdmins `field:"optional" json:"effectiveRestrictWorkspaceAdmins" yaml:"effectiveRestrictWorkspaceAdmins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#integer_val DataDatabricksWorkspaceSettingV2#integer_val}.
	IntegerVal *DataDatabricksWorkspaceSettingV2IntegerVal `field:"optional" json:"integerVal" yaml:"integerVal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#name DataDatabricksWorkspaceSettingV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#personal_compute DataDatabricksWorkspaceSettingV2#personal_compute}.
	PersonalCompute *DataDatabricksWorkspaceSettingV2PersonalCompute `field:"optional" json:"personalCompute" yaml:"personalCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#restrict_workspace_admins DataDatabricksWorkspaceSettingV2#restrict_workspace_admins}.
	RestrictWorkspaceAdmins *DataDatabricksWorkspaceSettingV2RestrictWorkspaceAdmins `field:"optional" json:"restrictWorkspaceAdmins" yaml:"restrictWorkspaceAdmins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#string_val DataDatabricksWorkspaceSettingV2#string_val}.
	StringVal *DataDatabricksWorkspaceSettingV2StringVal `field:"optional" json:"stringVal" yaml:"stringVal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2#workspace_id DataDatabricksWorkspaceSettingV2#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

