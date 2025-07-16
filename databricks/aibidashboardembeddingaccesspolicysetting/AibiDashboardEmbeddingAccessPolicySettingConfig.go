// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aibidashboardembeddingaccesspolicysetting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AibiDashboardEmbeddingAccessPolicySettingConfig struct {
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
	// aibi_dashboard_embedding_access_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/aibi_dashboard_embedding_access_policy_setting#aibi_dashboard_embedding_access_policy AibiDashboardEmbeddingAccessPolicySetting#aibi_dashboard_embedding_access_policy}
	AibiDashboardEmbeddingAccessPolicy *AibiDashboardEmbeddingAccessPolicySettingAibiDashboardEmbeddingAccessPolicy `field:"required" json:"aibiDashboardEmbeddingAccessPolicy" yaml:"aibiDashboardEmbeddingAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/aibi_dashboard_embedding_access_policy_setting#etag AibiDashboardEmbeddingAccessPolicySetting#etag}.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/aibi_dashboard_embedding_access_policy_setting#id AibiDashboardEmbeddingAccessPolicySetting#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/aibi_dashboard_embedding_access_policy_setting#setting_name AibiDashboardEmbeddingAccessPolicySetting#setting_name}.
	SettingName *string `field:"optional" json:"settingName" yaml:"settingName"`
}

