// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusterpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/cluster_policy#name ClusterPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/cluster_policy#definition ClusterPolicy#definition}.
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/cluster_policy#description ClusterPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/cluster_policy#id ClusterPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// libraries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/cluster_policy#libraries ClusterPolicy#libraries}
	Libraries interface{} `field:"optional" json:"libraries" yaml:"libraries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/cluster_policy#max_clusters_per_user ClusterPolicy#max_clusters_per_user}.
	MaxClustersPerUser *float64 `field:"optional" json:"maxClustersPerUser" yaml:"maxClustersPerUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/cluster_policy#policy_family_definition_overrides ClusterPolicy#policy_family_definition_overrides}.
	PolicyFamilyDefinitionOverrides *string `field:"optional" json:"policyFamilyDefinitionOverrides" yaml:"policyFamilyDefinitionOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/cluster_policy#policy_family_id ClusterPolicy#policy_family_id}.
	PolicyFamilyId *string `field:"optional" json:"policyFamilyId" yaml:"policyFamilyId"`
}

